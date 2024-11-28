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

type WhereChat struct {
	ID            string     `json:"id,omitempty"`
	RemoteJID     string     `json:"remoteJid,omitempty"`
	PushName      *string    `json:"pushName,omitempty,omitempty"`
	ProfilePicURL *string    `json:"profilePicUrl,omitempty,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty,omitempty"`
	InstanceID    string     `json:"instanceId,omitempty"`
}

type FindChatsRequest QueryFilters[WhereChat]

type FindChatsResponse struct {
	Id             string      `json:"id,omitempty"`
	RemoteJid      string      `json:"remoteJid,omitempty"`
	Name           any         `json:"name,omitempty"`
	Labels         []string    `json:"labels,omitempty"`
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	PushName       string      `json:"pushName,omitempty"`
	ProfilePicUrl  string      `json:"profilePicUrl,omitempty"`
	UnreadMessages int         `json:"unreadMessages,omitempty"`
	LastMessage    LastMessage `json:"lastMessage,omitempty"`
}

type LastMessage struct {
	Id               string                              `json:"id"`
	Key              FindChatsResponseLastMessageKey     `json:"key"`
	PushName         string                              `json:"pushName"`
	Participant      any                                 `json:"participant"`
	MessageType      string                              `json:"messageType"`
	Message          FindChatsResponseLastMessageMessage `json:"message"`
	ContextInfo      any                                 `json:"contextInfo"`
	Source           string                              `json:"source"`
	MessageTimestamp int                                 `json:"messageTimestamp"`
	InstanceId       string                              `json:"instanceId"`
	SessionId        any                                 `json:"sessionId"`
	Status           string                              `json:"status"`
}

type FindChatsResponseLastMessageKey struct {
	Id        string `json:"id"`
	FromMe    bool   `json:"fromMe"`
	RemoteJid string `json:"remoteJid"`
}

type FindChatsResponseLastMessageMessage struct {
	ImageMessage FindChatsResponseLastMessageMessageImageMessage `json:"imageMessage"`
}

type FindChatsResponseLastMessageMessageImageMessage struct {
	Url               string `json:"url"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	Caption           string `json:"caption"`
	MediaKey          string `json:"mediaKey"`
	Mimetype          string `json:"mimetype"`
	DirectPath        string `json:"directPath"`
	FileLength        string `json:"fileLength"`
	FileSha256        string `json:"fileSha256"`
	FileEncSha256     string `json:"fileEncSha256"`
	MediaKeyTimestamp string `json:"mediaKeyTimestamp"`
}

func (s *Client) FindChats(ctx context.Context, instanceName string, filter *FindChatsRequest) ([]FindChatsResponse, error) {
	resp, err := s.request(ctx, filter, http.MethodPost, fmt.Sprintf(findChatsEndpoint, instanceName))
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
		return nil, fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn []FindChatsResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return toReturn, nil
}

type ReadMessagesRequest struct {
	ReadMessages []ReadMessagesRequestItem `json:"readMessages"`
}
type ReadMessagesRequestItem struct {
	RemoteJid string `json:"remoteJid"`
	FromMe    bool   `json:"fromMe"`
	Id        string `json:"id"`
}

type ReadMessagesResponse struct {
	Message string `json:"message"`
	Read    string `json:"read"`
}

func (s *Client) ReadMessages(ctx context.Context, instanceName string, req *ReadMessagesRequest) (*ReadMessagesResponse, error) {
	url := fmt.Sprintf(readMessagesEndpoint, instanceName)

	if req == nil {
		return nil, fmt.Errorf("missing request object")
	}

	resp, err := s.request(ctx, req, http.MethodPost, url)
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
		return nil, fmt.Errorf("failed to read message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn ReadMessagesResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

type WhereMessage struct {
	RemoteJID  string `json:"remoteJid,omitempty"`
	ID         string `json:"id,omitempty"`
	FromMe     bool   `json:"fromMe,omitempty"`
	InstanceId string `json:"instanceId,omitempty"`
}

type FindMessagesRequest QueryFilters[WhereMessage]

type FindMessagesResponse struct {
	Messages MessagesInfo `json:"messages"`
}

type MessagesInfo struct {
	Total       int             `json:"total"`
	Pages       int             `json:"pages"`
	CurrentPage int             `json:"currentPage"`
	Records     []MessageRecord `json:"records"`
}

type MessageRecord struct {
	Id               string          `json:"id"`
	Key              MessageKey      `json:"key"`
	PushName         string          `json:"pushName"`
	MessageType      string          `json:"messageType"`
	Message          MessageDetail   `json:"message"`
	MessageTimestamp int             `json:"messageTimestamp"`
	InstanceId       string          `json:"instanceId"`
	Source           string          `json:"source"`
	MessageUpdate    []MessageUpdate `json:"MessageUpdate"`
}

type MessageKey struct {
	Id        string `json:"id"`
	FromMe    bool   `json:"fromMe"`
	RemoteJid string `json:"remoteJid"`
}

type MessageDetail struct {
	Conversation       string              `json:"conversation"`
	MessageContextInfo *MessageContextInfo `json:"messageContextInfo,omitempty"`
}

type MessageContextInfo struct {
	MessageSecret             string              `json:"messageSecret"`
	DeviceListMetadata        *DeviceListMetadata `json:"deviceListMetadata,omitempty"`
	DeviceListMetadataVersion int                 `json:"deviceListMetadataVersion,omitempty"`
}

type DeviceListMetadata struct {
	SenderKeyHash      string `json:"senderKeyHash"`
	SenderTimestamp    string `json:"senderTimestamp"`
	RecipientKeyHash   string `json:"recipientKeyHash"`
	RecipientTimestamp string `json:"recipientTimestamp"`
}

type MessageUpdate struct {
	Status string `json:"status"`
}

func (s *Client) FindMessages(ctx context.Context, instanceName string, req *FindMessagesRequest) (*FindMessagesResponse, error) {
	resp, err := s.request(ctx, req, http.MethodPost, fmt.Sprintf(findMessagesEndpoint, instanceName))
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
		return nil, fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn FindMessagesResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return &toReturn, nil
}

type UnreadChatRequest struct {
	LastMessage UnreadChatLastMessage `json:"lastMessage"`
	Chat        string                `json:"chat"`
}

type UnreadChatLastMessage struct {
	Key UnreadChatLastMessageKey `json:"key"`
}
type UnreadChatLastMessageKey struct {
	Id        string `json:"id"`
	FromMe    bool   `json:"fromMe"`
	RemoteJid string `json:"remoteJid"`
}
type UnreadChatResponse struct {
	ChatId           string `json:"chatId"`
	MarkedChatUnread bool   `json:"markedChatUnread"`
}

func (s *Client) UnreadChat(ctx context.Context, instanceName string, req *UnreadChatRequest) (*UnreadChatResponse, error) {
	resp, err := s.request(ctx, req, http.MethodPost, fmt.Sprintf(unreadChatEndpoint, instanceName))
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
		return nil, fmt.Errorf("failed to get qr code with code %d: %w", resp.StatusCode, bodyErr)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var toReturn UnreadChatResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return &toReturn, nil
}
