package evolution

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type QueryFilters[T any] struct {
	Where  T      `json:"where,omitempty"`
	Sort   string `json:"sort,omitempty"`
	Page   int    `json:"page,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

func (s *Client) request(ctx context.Context, reqBody any, method, endpoint string) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		marshalledBody, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(marshalledBody)
	}

	url := fmt.Sprintf("%s/%s", s.baseUrl, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", s.apiKey)

	return s.httpClient.Do(req)
}
