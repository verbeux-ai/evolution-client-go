package evolution_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/evolution-client-go"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendTextMessage(ctx, &evolution.TextMessageRequest{
		Number: "5585999999999",
		Text:   "Testing",
	}, "Testing")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
