package evolution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CreateInstanceRequest struct {
	InstanceName                    string                        `json:"instanceName"`
	Qrcode                          bool                          `json:"qrcode"`
	Integration                     string                        `json:"integration"`
	RejectCall                      bool                          `json:"rejectCall"`
	MsgCall                         string                        `json:"msgCall"`
	GroupsIgnore                    bool                          `json:"groupsIgnore"`
	AlwaysOnline                    bool                          `json:"alwaysOnline"`
	ReadMessages                    bool                          `json:"readMessages"`
	ReadStatus                      bool                          `json:"readStatus"`
	SyncFullHistory                 bool                          `json:"syncFullHistory"`
	ProxyHost                       string                        `json:"proxyHost"`
	ProxyPort                       string                        `json:"proxyPort"`
	ProxyProtocol                   string                        `json:"proxyProtocol"`
	ProxyUsername                   string                        `json:"proxyUsername"`
	ProxyPassword                   string                        `json:"proxyPassword"`
	Webhook                         CreateInstanceRequestWebhook  `json:"webhook"`
	Rabbitmq                        CreateInstanceRequestRabbitMQ `json:"rabbitmq"`
	Sqs                             CreateInstanceRequestSqs      `json:"sqs"`
	ChatwootAccountId               string                        `json:"chatwootAccountId"`
	ChatwootToken                   string                        `json:"chatwootToken"`
	ChatwootUrl                     string                        `json:"chatwootUrl"`
	ChatwootSignMsg                 bool                          `json:"chatwootSignMsg"`
	ChatwootReopenConversation      bool                          `json:"chatwootReopenConversation"`
	ChatwootConversationPending     bool                          `json:"chatwootConversationPending"`
	ChatwootImportContacts          bool                          `json:"chatwootImportContacts"`
	ChatwootNameInbox               string                        `json:"chatwootNameInbox"`
	ChatwootMergeBrazilContacts     bool                          `json:"chatwootMergeBrazilContacts"`
	ChatwootImportMessages          bool                          `json:"chatwootImportMessages"`
	ChatwootDaysLimitImportMessages int                           `json:"chatwootDaysLimitImportMessages"`
	ChatwootOrganization            string                        `json:"chatwootOrganization"`
	ChatwootLogo                    string                        `json:"chatwootLogo"`
}

type CreateInstanceRequestWebhook struct {
	Url      string            `json:"url"`
	ByEvents bool              `json:"byEvents"`
	Base64   bool              `json:"base64"`
	Headers  map[string]string `json:"headers"`
	Events   []string          `json:"events"`
}

type CreateInstanceRequestRabbitMQ struct {
	Enabled bool     `json:"enabled"`
	Events  []string `json:"events"`
}

type CreateInstanceRequestSqs struct {
	Enabled bool     `json:"enabled"`
	Events  []string `json:"events"`
}

type CreateInstanceResponse struct {
	Instance  Instance                       `json:"instance"`
	Hash      string                         `json:"hash"`
	Webhook   CreateInstanceWebhookResponse  `json:"webhook"`
	Websocket any                            `json:"websocket"`
	Rabbitmq  CreateInstanceRabbitMQResponse `json:"rabbitmq"`
	Sqs       CreateInstanceSqsResponse      `json:"sqs"`
	Settings  CreateInstanceSettingsResponse `json:"settings"`
	Qrcode    CreateInstanceQrCodeResponse   `json:"qrcode"`
}

type Instance struct {
	InstanceName          string      `json:"instanceName"`
	InstanceId            string      `json:"instanceId"`
	Integration           string      `json:"integration"`
	WebhookWaBusiness     interface{} `json:"webhookWaBusiness"`
	AccessTokenWaBusiness string      `json:"accessTokenWaBusiness"`
	Status                string      `json:"status"`
}

type CreateInstanceWebhookResponse struct {
	WebhookUrl      string            `json:"webhookUrl"`
	WebhookHeaders  map[string]string `json:"webhookHeaders"`
	WebhookByEvents bool              `json:"webhookByEvents"`
	WebhookBase64   bool              `json:"webhookBase64"`
}

type CreateInstanceRabbitMQResponse struct {
	Enabled bool `json:"enabled"`
}

type CreateInstanceSqsResponse struct {
	Enabled bool `json:"enabled"`
}

type CreateInstanceSettingsResponse struct {
	RejectCall      bool   `json:"rejectCall"`
	MsgCall         string `json:"msgCall"`
	GroupsIgnore    bool   `json:"groupsIgnore"`
	AlwaysOnline    bool   `json:"alwaysOnline"`
	ReadMessages    bool   `json:"readMessages"`
	ReadStatus      bool   `json:"readStatus"`
	SyncFullHistory bool   `json:"syncFullHistory"`
}

type CreateInstanceQrCodeResponse struct {
	PairingCode interface{} `json:"pairingCode"`
	Code        string      `json:"code"`
	Base64      string      `json:"base64"`
	Count       int         `json:"count"`
}

func (s *Client) CreateInstance(ctx context.Context, d *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	resp, err := s.request(ctx, d, http.MethodPost, createInstanceEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyErr := errors.New(string(body))
		return nil, fmt.Errorf("failed to start sessionKey with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn CreateInstanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

type RestartInstanceRequest struct {
	InstanceName string `json:"instanceName"`
}

type RestartInstanceResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (s *Client) RestartInstance(ctx context.Context, d *Instance) (*RestartInstanceResponse, error) {
	url := fmt.Sprintf("%s/%s", restartInstanceEndpoint, d.InstanceName)

	resp, err := s.request(ctx, d, http.MethodPost, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", readErr)
		}
		return nil, fmt.Errorf("failed to restart instance with code %d: %s", resp.StatusCode, string(body))
	}

	var result RestartInstanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

type LogoutInstanceResponse struct {
	Status   string         `json:"status"`
	Error    bool           `json:"error"`
	Response LogoutResponse `json:"response"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

func (s *Client) LogoutInstance(ctx context.Context, d *Instance) (*LogoutInstanceResponse, error) {
	url := fmt.Sprintf("%s/%s", logoutInstanceEndpoint, d.InstanceName)

	resp, err := s.request(ctx, d, http.MethodDelete, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", readErr)
		}
		return nil, fmt.Errorf("failed to logout instance with code %d: %s", resp.StatusCode, string(body))
	}

	var result LogoutInstanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

type DeleteInstanceResponse struct {
	Status   string         `json:"status"`
	Error    bool           `json:"error"`
	Response DeleteResponse `json:"response"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

func (s *Client) DeleteInstance(ctx context.Context, d *Instance) (*DeleteInstanceResponse, error) {
	url := fmt.Sprintf("%s/%s", deleteInstanceEndpoint, d.InstanceName)

	resp, err := s.request(ctx, d, http.MethodDelete, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", readErr)
		}
		return nil, fmt.Errorf("failed to delete instance with code %d: %s", resp.StatusCode, string(body))
	}

	var result DeleteInstanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

type GetInstanceConnectResponse struct {
	PairingCode interface{} `json:"pairingCode"`
	Code        string      `json:"code"`
	Base64      string      `json:"base64"`
	Count       int         `json:"count"`
}

func (s *Client) GetConnectInstance(ctx context.Context, d *Instance) (*GetInstanceConnectResponse, error) {
	url := fmt.Sprintf("%s/%s", getConnectInstanceEndpoint, d.InstanceName)

	resp, err := s.request(ctx, d, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", readErr)
		}
		return nil, fmt.Errorf("failed to get connect instance with code %d: %s", resp.StatusCode, string(body))
	}

	var result GetInstanceConnectResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

type GetConnectionStateInstanceResponse struct {
	Instance InstanceConnectionState `json:"instance"`
}

type InstanceConnectionState struct {
	InstanceName string `json:"instanceName"`
	State        string `json:"state"`
}

func (s *Client) GetConnectionStateInstance(ctx context.Context, d *Instance) (*GetConnectionStateInstanceResponse, error) {
	url := fmt.Sprintf("%s/%s", getConnectionStateInstanceEndpoint, d.InstanceName)

	resp, err := s.request(ctx, d, http.MethodGet, url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read error response body: %w", readErr)
		}
		return nil, fmt.Errorf("failed to get connection state from instance with code %d: %s", resp.StatusCode, string(body))
	}

	var result GetConnectionStateInstanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
