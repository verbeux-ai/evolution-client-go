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
	Key              MessageResponseKey             `json:"key"`
	PushName         string                         `json:"pushName"`
	Status           string                         `json:"status"`
	Message          TextMessageResponseMessage     `json:"message"`
	ContextInfo      TextMessageResponseContextInfo `json:"contextInfo"`
	MessageType      string                         `json:"messageType"`
	MessageTimestamp int                            `json:"messageTimestamp"`
	InstanceId       string                         `json:"instanceId"`
	Source           string                         `json:"source"`
}

type MessageResponseKey struct {
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

func (s *Client) SendTextMessage(ctx context.Context, instanceName string, req *TextMessageRequest) (*TextMessageResponse, error) {
	url := fmt.Sprintf(sendMessageTextEndpoint, instanceName)

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
	Number    string `json:"number,omitempty"`
	Mediatype string `json:"mediatype,omitempty"`
	Mimetype  string `json:"mimetype,omitempty"`
	Caption   string `json:"caption,omitempty"`
	// Media is the URL of the file
	Media            string                `json:"media,omitempty"`
	FileName         string                `json:"fileName,omitempty"`
	Delay            int                   `json:"delay,omitempty"`
	Quoted           *MessageRequestQuoted `json:"quoted,omitempty"`
	MentionsEveryOne bool                  `json:"mentionsEveryOne,omitempty"`
	Mentioned        []string              `json:"mentioned,omitempty"`
}

type MediaMessageResponse struct {
	Key              MessageResponseKey          `json:"key,omitempty"`
	PushName         string                      `json:"pushName,omitempty"`
	Status           string                      `json:"status,omitempty"`
	Message          MediaMessageResponseMessage `json:"message,omitempty"`
	ContextInfo      any                         `json:"contextInfo,omitempty"`
	MessageType      string                      `json:"messageType,omitempty"`
	MessageTimestamp int                         `json:"messageTimestamp,omitempty"`
	InstanceId       string                      `json:"instanceId,omitempty"`
	Source           string                      `json:"source,omitempty"`
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

func (s *Client) SendMediaMessage(ctx context.Context, instanceName string, req *MediaMessageRequest) (*TextMessageResponse, error) {
	url := fmt.Sprintf(sendMessageMediaEndpoint, instanceName)

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

type AudioMessageRequest struct {
	Number           string                `json:"number,omitempty"`
	Audio            string                `json:"audio,omitempty"`
	Delay            int                   `json:"delay,omitempty"`
	Quoted           *MessageRequestQuoted `json:"quoted,omitempty"`
	MentionsEveryOne bool                  `json:"mentionsEveryOne,omitempty"`
	Mentioned        []string              `json:"mentioned,omitempty"`
	Encoding         bool                  `json:"encoding,omitempty"`
}

type AudioMessageResponseMessage struct {
	AudioMessage AudioMessageResponseMessageAudio `json:"audioMessage"`
	Base64       string                           `json:"base64"`
}
type AudioMessageResponseMessageAudio struct {
	DirectPath        string `json:"directPath"`
	FileEncSha256     string `json:"fileEncSha256"`
	FileLength        string `json:"fileLength"`
	FileSha256        string `json:"fileSha256"`
	MediaKey          string `json:"mediaKey"`
	MediaKeyTimestamp string `json:"mediaKeyTimestamp"`
	Mimetype          string `json:"mimetype"`
	Ptt               bool   `json:"ptt"`
	Seconds           int    `json:"seconds"`
	Url               string `json:"url"`
	Waveform          string `json:"waveform"`
}

type AudioMessageResponse struct {
	ContextInfo      MessageContextInfo          `json:"contextInfo"`
	InstanceId       string                      `json:"instanceId"`
	Key              MessageResponseKey          `json:"key"`
	Message          AudioMessageResponseMessage `json:"message"`
	MessageTimestamp int                         `json:"messageTimestamp"`
	MessageType      string                      `json:"messageType"`
	PushName         string                      `json:"pushName"`
	Source           string                      `json:"source"`
	Status           string                      `json:"status"`
}

func (s *Client) SendAudioMessage(ctx context.Context, instanceName string, req *AudioMessageRequest) (*AudioMessageResponse, error) {
	url := fmt.Sprintf(sendMessageAudioEndpoint, instanceName)
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

	var toReturn AudioMessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}

type ListMessageRequest struct {
	Number           string                      `json:"number"`
	Title            string                      `json:"title"`
	Description      string                      `json:"description"`
	ButtonText       string                      `json:"buttonText"`
	FooterText       string                      `json:"footerText"`
	Sections         []ListMessageRequestSection `json:"sections"`
	Delay            *int                        `json:"delay,omitempty"`
	Quoted           *MessageRequestQuoted       `json:"quoted,omitempty"`
	MentionsEveryone bool                        `json:"mentionsEveryOne,omitempty"`
	Mentioned        []string                    `json:"mentioned,omitempty"`
}

type ListMessageRequestSection struct {
	Title string                         `json:"title"`
	Rows  []ListMessageRequestSectionRow `json:"rows"`
}

type ListMessageRequestSectionRow struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	RowID       string `json:"rowId"`
}

type ListMessageResponse interface{}

func (s *Client) SendListMessage(ctx context.Context, instanceName string, req *ListMessageRequest) (*ListMessageResponse, error) {
	url := fmt.Sprintf(sendMessageListEndpoint, instanceName)
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

	var toReturn ListMessageResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
