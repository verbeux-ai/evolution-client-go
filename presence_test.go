package evolution_test

import (
	"context"
	"os"
	"testing"

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
