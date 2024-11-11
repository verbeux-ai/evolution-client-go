package listener

import (
	"encoding/json"
	"io"
)

func (s *listener) ReadBodyAsync(rawBody io.ReadCloser) error {
	var data MessageUpsert
	if err := json.NewDecoder(rawBody).Decode(&data); err != nil {
		return err
	}

	if err := (*s.messageUpsertListener)(&data); err != nil {
		s.chError <- err
	}

	return nil
}
