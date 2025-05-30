package listener_test

import (
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/evolution-client-go/listener"
)

func TestListener_OnMessage(t *testing.T) {
	client := listener.NewEventListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnMessage(func(message *listener.MessageUpsert) error {
		defer wg.Done()
		require.NotEmpty(t, message)
		return nil
	})

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(messageExample1)))
	require.NoError(t, err)

	wg.Wait()

}

func TestListener_OnImageMessage(t *testing.T) {
	client := listener.NewEventListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnMessage(func(message *listener.MessageUpsert) error {
		defer wg.Done()
		require.NotEmpty(t, message)
		return nil
	})

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(messageImageExample1)))
	require.NoError(t, err)

	wg.Wait()

}

func TestListener_OnAudioMessage(t *testing.T) {
	client := listener.NewEventListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnMessage(func(message *listener.MessageUpsert) error {
		defer wg.Done()
		require.NotEmpty(t, message)
		return nil
	})

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(messageAudioExample1)))
	require.NoError(t, err)

	wg.Wait()

}

func TestListener_OnListMessage(t *testing.T) {
	client := listener.NewEventListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnMessage(func(message *listener.MessageUpsert) error {
		defer wg.Done()
		require.NotEmpty(t, message)
		require.NotEmpty(t, message.Data)
		require.NotEmpty(t, message.Data.Message)
		require.NotEmpty(t, message.Data.Message.ListResponseMessage)
		require.NotEmpty(t, message.Data.Message.ListResponseMessage.ListType)
		require.NotEmpty(t, message.Data.Message.ListResponseMessage.SingleSelectReply)
		require.NotEmpty(t, message.Data.Message.ListResponseMessage.SingleSelectReply.SelectedRowId)
		return nil
	})

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(messageListExample1)))
	require.NoError(t, err)

	wg.Wait()

}

func TestListener_OnContactUpsert(t *testing.T) {
	client := listener.NewEventListener()

	wg := sync.WaitGroup{}
	wg.Add(1)

	client.OnContactUpsert(func(contacts []listener.Contact) error {
		defer wg.Done()
		require.NotEmpty(t, contacts)
		require.Len(t, contacts, 1)
		require.NotEmpty(t, contacts[0].RemoteJid)
		require.NotEmpty(t, contacts[0].PushName)
		require.Nil(t, contacts[0].ProfilePicUrl)
		require.NotEmpty(t, contacts[0].InstanceId)
		return nil
	})

	err := client.ReadBodyAsync(io.NopCloser(strings.NewReader(contactUpsertExample1)))
	require.NoError(t, err)

	wg.Wait()
}
