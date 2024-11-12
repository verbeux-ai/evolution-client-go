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
	InstanceName                    string                         `json:"instanceName,omitempty"`
	Qrcode                          bool                           `json:"qrcode,omitempty"`
	Integration                     string                         `json:"integration,omitempty"`
	RejectCall                      bool                           `json:"rejectCall,omitempty"`
	MsgCall                         string                         `json:"msgCall,omitempty"`
	GroupsIgnore                    bool                           `json:"groupsIgnore,omitempty"`
	AlwaysOnline                    bool                           `json:"alwaysOnline,omitempty"`
	ReadMessages                    bool                           `json:"readMessages,omitempty"`
	ReadStatus                      bool                           `json:"readStatus,omitempty"`
	SyncFullHistory                 bool                           `json:"syncFullHistory,omitempty"`
	ProxyHost                       string                         `json:"proxyHost,omitempty"`
	ProxyPort                       string                         `json:"proxyPort,omitempty"`
	ProxyProtocol                   string                         `json:"proxyProtocol,omitempty"`
	ProxyUsername                   string                         `json:"proxyUsername,omitempty"`
	ProxyPassword                   string                         `json:"proxyPassword,omitempty"`
	Webhook                         *CreateInstanceRequestWebhook  `json:"webhook,omitempty"`
	Rabbitmq                        *CreateInstanceRequestRabbitMQ `json:"rabbitmq,omitempty"`
	Sqs                             *CreateInstanceRequestSqs      `json:"sqs,omitempty"`
	ChatwootAccountId               string                         `json:"chatwootAccountId,omitempty"`
	ChatwootToken                   string                         `json:"chatwootToken,omitempty"`
	ChatwootUrl                     string                         `json:"chatwootUrl,omitempty"`
	ChatwootSignMsg                 bool                           `json:"chatwootSignMsg,omitempty"`
	ChatwootReopenConversation      bool                           `json:"chatwootReopenConversation,omitempty"`
	ChatwootConversationPending     bool                           `json:"chatwootConversationPending,omitempty"`
	ChatwootImportContacts          bool                           `json:"chatwootImportContacts,omitempty"`
	ChatwootNameInbox               string                         `json:"chatwootNameInbox,omitempty"`
	ChatwootMergeBrazilContacts     bool                           `json:"chatwootMergeBrazilContacts,omitempty"`
	ChatwootImportMessages          bool                           `json:"chatwootImportMessages,omitempty"`
	ChatwootDaysLimitImportMessages int                            `json:"chatwootDaysLimitImportMessages,omitempty"`
	ChatwootOrganization            string                         `json:"chatwootOrganization,omitempty"`
	ChatwootLogo                    string                         `json:"chatwootLogo,omitempty"`
}

type CreateInstanceRequestWebhook struct {
	Url      string            `json:"url,omitempty"`
	ByEvents bool              `json:"byEvents,omitempty"`
	Base64   bool              `json:"base64,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
	Events   []string          `json:"events,omitempty"`
}

type CreateInstanceRequestRabbitMQ struct {
	Enabled bool     `json:"enabled,omitempty"`
	Events  []string `json:"events,omitempty"`
}

type CreateInstanceRequestSqs struct {
	Enabled bool     `json:"enabled,omitempty"`
	Events  []string `json:"events,omitempty"`
}

type CreateInstanceResponse struct {
	Instance  Instance                       `json:"instance,omitempty"`
	Hash      string                         `json:"hash,omitempty"`
	Webhook   CreateInstanceWebhookResponse  `json:"webhook,omitempty"`
	Websocket any                            `json:"websocket,omitempty"`
	Rabbitmq  CreateInstanceRabbitMQResponse `json:"rabbitmq,omitempty"`
	Sqs       CreateInstanceSqsResponse      `json:"sqs,omitempty"`
	Settings  CreateInstanceSettingsResponse `json:"settings,omitempty"`
	Qrcode    CreateInstanceQrCodeResponse   `json:"qrcode,omitempty"`
}

type Instance struct {
	InstanceName          string      `json:"instanceName,omitempty"`
	InstanceId            string      `json:"instanceId,omitempty"`
	Integration           string      `json:"integration,omitempty"`
	WebhookWaBusiness     interface{} `json:"webhookWaBusiness,omitempty"`
	AccessTokenWaBusiness string      `json:"accessTokenWaBusiness,omitempty"`
	Status                string      `json:"status,omitempty"`
}

type CreateInstanceWebhookResponse struct {
	WebhookUrl      string            `json:"webhookUrl,omitempty"`
	WebhookHeaders  map[string]string `json:"webhookHeaders,omitempty"`
	WebhookByEvents bool              `json:"webhookByEvents,omitempty"`
	WebhookBase64   bool              `json:"webhookBase64,omitempty"`
}

type CreateInstanceRabbitMQResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

type CreateInstanceSqsResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

type CreateInstanceSettingsResponse struct {
	RejectCall      bool   `json:"rejectCall,omitempty"`
	MsgCall         string `json:"msgCall,omitempty"`
	GroupsIgnore    bool   `json:"groupsIgnore,omitempty"`
	AlwaysOnline    bool   `json:"alwaysOnline,omitempty"`
	ReadMessages    bool   `json:"readMessages,omitempty"`
	ReadStatus      bool   `json:"readStatus,omitempty"`
	SyncFullHistory bool   `json:"syncFullHistory,omitempty"`
}

type CreateInstanceQrCodeResponse struct {
	PairingCode interface{} `json:"pairingCode,omitempty"`
	Code        string      `json:"code,omitempty"`
	Base64      string      `json:"base64,omitempty"`
	Count       int         `json:"count,omitempty"`
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

type RestartInstanceResponse struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
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
	Status   string         `json:"status,omitempty"`
	Error    bool           `json:"error,omitempty"`
	Response LogoutResponse `json:"response,omitempty"`
}

type LogoutResponse struct {
	Message string `json:"message,omitempty"`
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
	Status   string         `json:"status,omitempty"`
	Error    bool           `json:"error,omitempty"`
	Response DeleteResponse `json:"response,omitempty"`
}

type DeleteResponse struct {
	Message string `json:"message,omitempty"`
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
	PairingCode interface{} `json:"pairingCode,omitempty"`
	Code        string      `json:"code,omitempty"`
	Base64      string      `json:"base64,omitempty"`
	Count       int         `json:"count,omitempty"`
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
	Instance InstanceConnectionState `json:"instance,omitempty"`
}

type InstanceConnectionState struct {
	InstanceName string `json:"instanceName,omitempty"`
	State        string `json:"state,omitempty"`
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
