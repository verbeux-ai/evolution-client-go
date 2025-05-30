package listener

import (
	"io"
)

type listener struct {
	chError chan error

	messageUpsertListener  *MessageUpsertListener
	presenceUpdateListener *PresenceUpdateListener
	messageUpdateListener  *MessageUpdateListener
	contactUpdateListener  *ContactUpdateListener
	contactUpsertListener  *ContactUpsertListener
}

func NewEventListener() EventListener {
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

type EventListener interface {
	HandleErrors(f func(error)) (closer func())
	OnMessage(MessageUpsertListener)
	OnPresence(PresenceUpdateListener)
	OnMessageUpdate(MessageUpdateListener)
	OnContactUpdate(ContactUpdateListener)
	OnContactUpsert(ContactUpsertListener)
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

func (s *listener) OnContactUpdate(listener ContactUpdateListener) {
	s.contactUpdateListener = &listener
}

func (s *listener) OnContactUpsert(listener ContactUpsertListener) {
	s.contactUpsertListener = &listener
}
