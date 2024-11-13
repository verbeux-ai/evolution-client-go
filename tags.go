package evolution

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TagResponse struct {
	Color        string `json:"color"`
	Name         string `json:"name"`
	Id           string `json:"id"`
	PredefinedId string `json:"predefinedId"`
}

func (s *Client) GetTags(ctx context.Context, instanceName string) ([]TagResponse, error) {
	resp, err := s.request(ctx, nil, http.MethodGet, fmt.Sprintf(findLabelsEndpoint, instanceName))
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

	var toReturn []TagResponse
	if err = json.Unmarshal(body, &toReturn); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(body))
	}

	return toReturn, nil
}

type addChatTagRequestInternal struct {
	Number  string `json:"number"`
	LabelId string `json:"labelId"`
	Action  string `json:"action"`
}

func (s *Client) AddChatTag(ctx context.Context, instanceName, tagID, phone string) error {
	req := addChatTagRequestInternal{
		Number:  phone,
		LabelId: tagID,
		Action:  "add",
	}

	resp, err := s.request(ctx, req, http.MethodPost, fmt.Sprintf(handleLabelEndpoint, instanceName))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		bodyErr := errors.New(string(body))
		return fmt.Errorf("failed to set chat tag with code %d: %w", resp.StatusCode, bodyErr)
	}

	return nil
}
