package evolution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TextMessageRequest struct {
	Number           string                `json:"number,omitempty"`
	Text             string                `json:"text,omitempty"`
	Delay            int                   `json:"delay,omitempty"`
	Quoted           *MessageRequestQuoted `json:"quoted,omitempty"`
	LinkPreview      bool                  `json:"linkPreview,omitempty"`
	MentionsEveryOne bool                  `json:"mentionsEveryOne,omitempty"`
	Mentioned        []string              `json:"mentioned,omitempty"`
}

type MessageRequestQuoted struct {
	Key     QuotedKey     `json:"key,omitempty"`
	Message QuotedMessage `json:"message,omitempty"`
}

type QuotedKey struct {
	Id string `json:"id,omitempty"`
}

type QuotedMessage struct {
	Conversation string `json:"conversation,omitempty"`
}

type TextMessageResponse struct {
	Key              TextMessageResponseKey         `json:"key"`
	PushName         string                         `json:"pushName"`
	Status           string                         `json:"status"`
	Message          TextMessageResponseMessage     `json:"message"`
	ContextInfo      TextMessageResponseContextInfo `json:"contextInfo"`
	MessageType      string                         `json:"messageType"`
	MessageTimestamp int                            `json:"messageTimestamp"`
	InstanceId       string                         `json:"instanceId"`
	Source           string                         `json:"source"`
}

type TextMessageResponseKey struct {
	RemoteJid string `json:"remoteJid,omitempty"`
	FromMe    bool   `json:"fromMe,omitempty"`
	Id        string `json:"id,omitempty"`
}

type TextMessageResponseMessage struct {
	Conversation string `json:"conversation,omitempty"`
}

type TextMessageResponseContextInfo struct {
	Participant   string                   `json:"participant,omitempty"`
	StanzaId      string                   `json:"stanzaId,omitempty"`
	QuotedMessage ContextInfoQuotedMessage `json:"quotedMessage,omitempty"`
}

type ContextInfoQuotedMessage struct {
	Conversation string `json:"conversation,omitempty"`
}

func (s *Client) SendTextMessage(ctx context.Context, req *TextMessageRequest, instanceName string) (*TextMessageResponse, error) {
	url := fmt.Sprintf("%s/%s", sendMessageTextEndpoint, instanceName)

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
		return nil, fmt.Errorf("failed to send text message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn TextMessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

type MediaMessageRequest struct {
	Number           string                `json:"number,omitempty"`
	Mediatype        string                `json:"mediatype,omitempty"`
	Mimetype         string                `json:"mimetype,omitempty"`
	Caption          string                `json:"caption,omitempty"`
	Media            string                `json:"media,omitempty"`
	FileName         string                `json:"fileName,omitempty"`
	Delay            int                   `json:"delay,omitempty"`
	Quoted           *MessageRequestQuoted `json:"quoted,omitempty"`
	MentionsEveryOne bool                  `json:"mentionsEveryOne,omitempty"`
	Mentioned        []string              `json:"mentioned,omitempty"`
}

type MediaMessageResponse struct {
	Key              MediaMessageResponseKey     `json:"key,omitempty"`
	PushName         string                      `json:"pushName,omitempty"`
	Status           string                      `json:"status,omitempty"`
	Message          MediaMessageResponseMessage `json:"message,omitempty"`
	ContextInfo      any                         `json:"contextInfo,omitempty"`
	MessageType      string                      `json:"messageType,omitempty"`
	MessageTimestamp int                         `json:"messageTimestamp,omitempty"`
	InstanceId       string                      `json:"instanceId,omitempty"`
	Source           string                      `json:"source,omitempty"`
}

type MediaMessageResponseKey struct {
	RemoteJid string `json:"remoteJid,omitempty"`
	FromMe    bool   `json:"fromMe,omitempty"`
	Id        string `json:"id,omitempty"`
}

type MediaMessageResponseMessage struct {
	ImageMessage MediaMessageResponseMessageImage `json:"imageMessage,omitempty"`
	Base64       string                           `json:"base64,omitempty"`
}

type MediaMessageResponseMessageImage struct {
	Url               string `json:"url,omitempty"`
	Mimetype          string `json:"mimetype,omitempty"`
	Caption           string `json:"caption,omitempty"`
	FileSha256        string `json:"fileSha256,omitempty"`
	FileLength        string `json:"fileLength,omitempty"`
	Height            int    `json:"height,omitempty"`
	Width             int    `json:"width,omitempty"`
	MediaKey          string `json:"mediaKey,omitempty"`
	FileEncSha256     string `json:"fileEncSha256,omitempty"`
	DirectPath        string `json:"directPath,omitempty"`
	MediaKeyTimestamp string `json:"mediaKeyTimestamp,omitempty"`
	JpegThumbnail     string `json:"jpegThumbnail,omitempty"`
	ContextInfo       any    `json:"contextInfo,omitempty"`
}

func (s *Client) SendMediaMessage(ctx context.Context, req *MediaMessageRequest, instanceName string) (*TextMessageResponse, error) {
	url := fmt.Sprintf("%s/%s", sendMessageMediaEndpoint, instanceName)

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
		return nil, fmt.Errorf("failed to send text message with code %d: %w", resp.StatusCode, bodyErr)
	}

	var toReturn TextMessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
