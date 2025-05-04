package listener

import (
	"io"
)

type listener struct {
	chError chan error

	messageUpsertListener  *MessageUpsertListener
	presenceUpdateListener *PresenceUpdateListener
	messageUpdateListener  *MessageUpdateListener
}

func NewMessageListener() MessageListener {
	return &listener{
		chError: make(chan error),
	}
}

func (s *listener) HandleErrors(f func(error)) (closer func()) {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-s.chError:
				f(err)
			case <-done:
				return
			}

		}
	}()

	return func() {
		done <- struct{}{}
	}
}

type MessageListener interface {
	HandleErrors(f func(error)) (closer func())
	OnMessage(MessageUpsertListener)
	OnPresence(PresenceUpdateListener)
	OnMessageUpdate(MessageUpdateListener)
	ReadBodyAsync(rawBody io.ReadCloser) error
}

func (s *listener) OnMessage(message MessageUpsertListener) {
	s.messageUpsertListener = &message
}

func (s *listener) OnPresence(presence PresenceUpdateListener) {
	s.presenceUpdateListener = &presence
}

func (s *listener) OnMessageUpdate(message MessageUpdateListener) {
	s.messageUpdateListener = &message
}
