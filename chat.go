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
	Id             string                       `json:"id,omitempty"`
	RemoteJid      string                       `json:"remoteJid,omitempty"`
	Name           any                          `json:"name,omitempty"`
	Labels         []string                     `json:"labels,omitempty"`
	CreatedAt      time.Time                    `json:"createdAt,omitempty"`
	UpdatedAt      time.Time                    `json:"updatedAt,omitempty"`
	PushName       string                       `json:"pushName,omitempty"`
	ProfilePicUrl  string                       `json:"profilePicUrl,omitempty"`
	UnreadMessages int                          `json:"unreadMessages,omitempty"`
	LastMessage    FindChatsResponseLastMessage `json:"lastMessage,omitempty"`
}

type FindChatsResponseLastMessage struct {
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
