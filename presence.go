package evolution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type SendPresenceRequestPresence string

const (
	PresenceComposing SendPresenceRequestPresence = "composing"
	PresenceAvailable SendPresenceRequestPresence = "available"
)

type SendPresenceRequest struct {
	Number   string                      `json:"number"`
	Delay    int                         `json:"delay"`
	Presence SendPresenceRequestPresence `json:"presence"`
}

type SendPresenceResponse struct {
	Presence SendPresenceRequestPresence `json:"presence"`
}

func (s *Client) SendPresence(ctx context.Context, instanceName string, req *SendPresenceRequest) (*SendPresenceResponse, error) {
	url := fmt.Sprintf(sendPresenceEndpoint, instanceName)

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

	var toReturn SendPresenceResponse
	if err = json.NewDecoder(resp.Body).Decode(&toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
