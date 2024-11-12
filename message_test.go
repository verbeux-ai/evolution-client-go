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
		Delay:  0,
		Quoted: evolution.MessageRequestQuoted{
			Key:     evolution.QuotedKey{Id: "1"},
			Message: evolution.QuotedMessage{},
		},
		LinkPreview:      false,
		MentionsEveryOne: false,
		Mentioned:        nil,
	}, "Testing")
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
