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
	Number           string               `json:"number"`
	Text             string               `json:"text"`
	Delay            int                  `json:"delay"`
	Quoted           MessageRequestQuoted `json:"quoted"`
	LinkPreview      bool                 `json:"linkPreview"`
	MentionsEveryOne bool                 `json:"mentionsEveryOne"`
	Mentioned        []string             `json:"mentioned"`
}

type MessageRequestQuoted struct {
	Key     QuotedKey     `json:"key"`
	Message QuotedMessage `json:"message"`
}

type QuotedKey struct {
	Id string `json:"id"`
}

type QuotedMessage struct {
	Conversation string `json:"conversation"`
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
	RemoteJid string `json:"remoteJid"`
	FromMe    bool   `json:"fromMe"`
	Id        string `json:"id"`
}

type TextMessageResponseMessage struct {
	Conversation string `json:"conversation"`
}

type TextMessageResponseContextInfo struct {
	Participant   string                   `json:"participant"`
	StanzaId      string                   `json:"stanzaId"`
	QuotedMessage ContextInfoQuotedMessage `json:"quotedMessage"`
}

type ContextInfoQuotedMessage struct {
	Conversation string `json:"conversation"`
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
	Number           string               `json:"number"`
	Mediatype        string               `json:"mediatype"`
	Mimetype         string               `json:"mimetype"`
	Caption          string               `json:"caption"`
	Media            string               `json:"media"`
	FileName         string               `json:"fileName"`
	Delay            int                  `json:"delay"`
	Quoted           MessageRequestQuoted `json:"quoted"`
	MentionsEveryOne bool                 `json:"mentionsEveryOne"`
	Mentioned        []string             `json:"mentioned"`
}

type MediaMessageResponse struct {
	Key              MediaMessageResponseKey     `json:"key"`
	PushName         string                      `json:"pushName"`
	Status           string                      `json:"status"`
	Message          MediaMessageResponseMessage `json:"message"`
	ContextInfo      any                         `json:"contextInfo"`
	MessageType      string                      `json:"messageType"`
	MessageTimestamp int                         `json:"messageTimestamp"`
	InstanceId       string                      `json:"instanceId"`
	Source           string                      `json:"source"`
}

type MediaMessageResponseKey struct {
	RemoteJid string `json:"remoteJid"`
	FromMe    bool   `json:"fromMe"`
	Id        string `json:"id"`
}

type MediaMessageResponseMessage struct {
	ImageMessage MediaMessageResponseMessageImage `json:"imageMessage"`
	Base64       string                           `json:"base64"`
}

type MediaMessageResponseMessageImage struct {
	Url               string `json:"url"`
	Mimetype          string `json:"mimetype"`
	Caption           string `json:"caption"`
	FileSha256        string `json:"fileSha256"`
	FileLength        string `json:"fileLength"`
	Height            int    `json:"height"`
	Width             int    `json:"width"`
	MediaKey          string `json:"mediaKey"`
	FileEncSha256     string `json:"fileEncSha256"`
	DirectPath        string `json:"directPath"`
	MediaKeyTimestamp string `json:"mediaKeyTimestamp"`
	JpegThumbnail     string `json:"jpegThumbnail"`
	ContextInfo       any    `json:"contextInfo"`
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
