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
		Text:   "Fala Dotô!",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
