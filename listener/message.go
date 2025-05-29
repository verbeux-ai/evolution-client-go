package listener

import (
	"encoding/json"
	"io"
)

func (s *listener) ReadBodyAsync(rawBody io.ReadCloser) error {
	defer rawBody.Close()

	rawData, err := io.ReadAll(rawBody)
	if err != nil {
		return err
	}

	var identifier WookIdentifier
	if err := json.Unmarshal(rawData, &identifier); err != nil {
		return err
	}

	switch identifier.Event {
	case WookTypeMessageUpsert:
		return s.handleMessageUpsert(rawData)
	case WookTypePresenceUpdate:
		return s.handlePresenceUpdate(rawData)
	case WookTypeMessageUpdate:
		return s.handleMessageUpdate(rawData)
	case WookTypeContactUpdate:
		return s.handleContactUpdate(rawData)
	case WookTypeContactUpsert:
		return s.handleContactUpsert(rawData)
	}

	return nil
}

func (s *listener) handleMessageUpsert(rawData []byte) error {
	var msg MessageUpsert
	if err := json.Unmarshal(rawData, &msg); err != nil {
		return err
	}
	if err := (*s.messageUpsertListener)(&msg); err != nil {
		s.chError <- err
	}
	return nil
}

func (s *listener) handleMessageUpdate(rawData []byte) error {
	var msg MessageUpdate
	if err := json.Unmarshal(rawData, &msg); err != nil {
		return err
	}
	if err := (*s.messageUpdateListener)(&msg); err != nil {
		s.chError <- err
	}
	return nil
}

func (s *listener) handlePresenceUpdate(rawData []byte) error {
	var presence PresenceUpdate
	if err := json.Unmarshal(rawData, &presence); err != nil {
		return err
	}
	if err := (*s.presenceUpdateListener)(&presence); err != nil {
		s.chError <- err
	}
	return nil
}

func (s *listener) handleContactUpdate(rawData []byte) error {
	var contact Contact
	if err := json.Unmarshal(rawData, &contact); err != nil {
		return err
	}

	if err := (*s.contactUpdateListener)(&contact); err != nil {
		s.chError <- err
	}
	return nil
}

func (s *listener) handleContactUpsert(rawData []byte) error {
	var contact Contact
	if err := json.Unmarshal(rawData, &contact); err != nil {
		return err
	}

	if err := (*s.contactUpsertListener)(&contact); err != nil {
		s.chError <- err
	}
	return nil
}
