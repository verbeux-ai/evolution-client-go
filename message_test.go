package evolution_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/evolution-client-go"
)

func TestSendTextMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTextMessage(ctx, os.Getenv("INSTANCE_NAME"), &evolution.TextMessageRequest{
		Number: os.Getenv("NUMBER"),
		Text:   "Teste",
		Delay:  5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestSendAudioMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendAudioMessage(ctx, os.Getenv("INSTANCE_NAME"), &evolution.AudioMessageRequest{
		Number: os.Getenv("NUMBER"),
		Audio:  "",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
