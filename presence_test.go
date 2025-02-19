package evolution_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/evolution-client-go"
)

func TestSendPresenceComposing(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendPresence(ctx, os.Getenv("INSTANCE_NAME"), &evolution.SendPresenceRequest{
		Number:   os.Getenv("NUMBER"),
		Delay:    1000,
		Presence: evolution.PresenceComposing,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.Presence, evolution.PresenceComposing)
}

func TestSendPresenceAvailable(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendPresence(ctx, os.Getenv("INSTANCE_NAME"), &evolution.SendPresenceRequest{
		Number:   os.Getenv("NUMBER"),
		Delay:    1000,
		Presence: evolution.PresenceAvailable,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.Presence, evolution.PresenceAvailable)
}

func TestSendPresenceAvailableAsync(t *testing.T) {
	ctx := context.Background()
	start := time.Now()
	result := client.SendPresenceAsync(ctx, os.Getenv("INSTANCE_NAME"), &evolution.SendPresenceRequest{
		Number:   os.Getenv("NUMBER"),
		Delay:    5000,
		Presence: evolution.PresenceAvailable,
	})
	require.Less(t, time.Since(start).Milliseconds(), time.Millisecond*1000, "Request should be asynchronous")

	select {
	case err := <-result:
		require.NoError(t, err, "Expected no error from SendPresenceAsync")
	case <-time.After(6 * time.Second):
		t.Fatal("Timeout waiting for SendPresenceAsync response")
	}
}
