package evolution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
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
	SyncRecentHistory               bool                           `json:"syncRecentHistory,omitempty"`
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
	InstanceName          string `json:"instanceName,omitempty"`
	InstanceId            string `json:"instanceId,omitempty"`
	Integration           string `json:"integration,omitempty"`
	WebhookWaBusiness     any    `json:"webhookWaBusiness,omitempty"`
	AccessTokenWaBusiness string `json:"accessTokenWaBusiness,omitempty"`
	Status                string `json:"status,omitempty"`
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
	PairingCode any    `json:"pairingCode,omitempty"`
	Code        string `json:"code,omitempty"`
	Base64      string `json:"base64,omitempty"`
	Count       int    `json:"count,omitempty"`
}

func (s *Client) CreateInstance(ctx context.Context, req *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	resp, err := s.request(ctx, req, http.MethodPost, createInstanceEndpoint)
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
	Instance Instance `json:"instance"`
}

func (s *Client) RestartInstance(ctx context.Context, instanceName string) (*RestartInstanceResponse, error) {
	url := fmt.Sprintf(restartInstanceEndpoint, instanceName)

	resp, err := s.request(ctx, nil, http.MethodPost, url)
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

func (s *Client) LogoutInstance(ctx context.Context, instanceName string) (*LogoutInstanceResponse, error) {
	url := fmt.Sprintf(logoutInstanceEndpoint, instanceName)

	resp, err := s.request(ctx, nil, http.MethodDelete, url)
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

func (s *Client) DeleteInstance(ctx context.Context, instanceName string) (*DeleteInstanceResponse, error) {
	url := fmt.Sprintf(deleteInstanceEndpoint, instanceName)

	resp, err := s.request(ctx, nil, http.MethodDelete, url)
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
	PairingCode any    `json:"pairingCode,omitempty"`
	Code        string `json:"code,omitempty"`
	Base64      string `json:"base64,omitempty"`
	Count       int    `json:"count,omitempty"`
}

func (s *Client) ConnectInstance(ctx context.Context, instanceName string) (*GetInstanceConnectResponse, error) {
	url := fmt.Sprintf(getConnectInstanceEndpoint, instanceName)

	resp, err := s.request(ctx, nil, http.MethodGet, url)
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

func (s *Client) ConnectionStateInstance(ctx context.Context, instanceName string) (*GetConnectionStateInstanceResponse, error) {
	url := fmt.Sprintf(getConnectionStateInstanceEndpoint, instanceName)

	resp, err := s.request(ctx, nil, http.MethodGet, url)
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

type FetchInstancesRequestFilter struct {
	InstanceName string `query:"instanceName"`
	InstanceID   string `query:"instanceId,omitempty"`
	QRCode       bool   `query:"qrcode,omitempty"`
	BusinessID   string `query:"businessId,omitempty"`
	Number       string `query:"number,omitempty"`
	Integration  string `query:"integration,omitempty"`
	Token        string `query:"token,omitempty"`
	Status       string `query:"status,omitempty"`

	// Settings
	RejectCall      bool   `query:"rejectCall,omitempty"`
	MsgCall         string `query:"msgCall,omitempty"`
	GroupsIgnore    bool   `query:"groupsIgnore,omitempty"`
	AlwaysOnline    bool   `query:"alwaysOnline,omitempty"`
	ReadMessages    bool   `query:"readMessages,omitempty"`
	ReadStatus      bool   `query:"readStatus,omitempty"`
	SyncFullHistory bool   `query:"syncFullHistory,omitempty"`

	// Proxy
	ProxyHost     string `query:"proxyHost,omitempty"`
	ProxyPort     string `query:"proxyPort,omitempty"`
	ProxyProtocol string `query:"proxyProtocol,omitempty"`
	ProxyUsername string `query:"proxyUsername,omitempty"`
	ProxyPassword string `query:"proxyPassword,omitempty"`
}

type FetchInstancesResponse struct {
	Id                      string                         `json:"id"`
	Name                    string                         `json:"name"`
	ConnectionStatus        string                         `json:"connectionStatus"`
	OwnerJid                string                         `json:"ownerJid"`
	ProfileName             string                         `json:"profileName"`
	ProfilePicUrl           string                         `json:"profilePicUrl"`
	Integration             string                         `json:"integration"`
	Number                  any                            `json:"number"`
	BusinessId              any                            `json:"businessId"`
	Token                   string                         `json:"token"`
	ClientName              string                         `json:"clientName"`
	DisconnectionReasonCode any                            `json:"disconnectionReasonCode"`
	DisconnectionObject     any                            `json:"disconnectionObject"`
	DisconnectionAt         any                            `json:"disconnectionAt"`
	CreatedAt               time.Time                      `json:"createdAt"`
	UpdatedAt               time.Time                      `json:"updatedAt"`
	Chatwoot                any                            `json:"Chatwoot"`
	Proxy                   any                            `json:"Proxy"`
	Rabbitmq                any                            `json:"Rabbitmq"`
	Sqs                     any                            `json:"Sqs"`
	Websocket               any                            `json:"Websocket"`
	Setting                 CreateInstanceResponseSettings `json:"Setting"`
	Count                   FetchInstancesResponseCount    `json:"_count"`
}

type FetchInstancesResponseCount struct {
	Message int `json:"Message"`
	Contact int `json:"Contact"`
	Chat    int `json:"Chat"`
}

type CreateInstanceResponseSettings struct {
	Id              string    `json:"id"`
	RejectCall      bool      `json:"rejectCall"`
	MsgCall         string    `json:"msgCall"`
	GroupsIgnore    bool      `json:"groupsIgnore"`
	AlwaysOnline    bool      `json:"alwaysOnline"`
	ReadMessages    bool      `json:"readMessages"`
	ReadStatus      bool      `json:"readStatus"`
	SyncFullHistory bool      `json:"syncFullHistory"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	InstanceId      string    `json:"instanceId"`
}

func (s *Client) FetchInstances(ctx context.Context, filter *FetchInstancesRequestFilter) ([]FetchInstancesResponse, error) {
	var query string
	var err error

	if filter != nil {
		query, err = StructToQueryString(*filter)
		if err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%s?%s", fetchInstancesEndpoint, query)
	resp, err := s.request(ctx, nil, http.MethodGet, url)
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

	var result []FetchInstancesResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}
