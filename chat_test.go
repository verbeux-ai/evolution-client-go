package evolution_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/verbeux-ai/evolution-client-go"
)

func TestFindChats(t *testing.T) {
	ctx := context.Background()
	result, err := client.FindChats(ctx, os.Getenv("INSTANCE_NAME"), &evolution.FindChatsRequest{
		Where: evolution.WhereChat{
			RemoteJID: os.Getenv("NUMBER"),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestFindMessages(t *testing.T) {
	ctx := context.Background()
	result, err := client.FindMessages(ctx, os.Getenv("INSTANCE_NAME"), &evolution.FindMessagesRequest{
		Where: evolution.WhereMessage{
			RemoteJID: os.Getenv("NUMBER"),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestMarkUnread(t *testing.T) {
	ctx := context.Background()
	result, err := client.FindMessages(ctx, os.Getenv("INSTANCE_NAME"), &evolution.FindMessagesRequest{
		Where: evolution.WhereMessage{
			RemoteJID: os.Getenv("NUMBER"),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Greater(t, result.Messages.Total, 0)
	require.Greater(t, len(result.Messages.Records), 0)

	resultUnread, err := client.UnreadChat(ctx, os.Getenv("INSTANCE_NAME"), &evolution.UnreadChatRequest{
		LastMessage: evolution.UnreadChatLastMessage{
			Key: evolution.UnreadChatLastMessageKey{
				Id:        result.Messages.Records[0].Key.Id,
				FromMe:    result.Messages.Records[0].Key.FromMe,
				RemoteJid: result.Messages.Records[0].Key.RemoteJid,
			},
		},
		Chat: result.Messages.Records[0].Key.RemoteJid,
	})
	require.NoError(t, err)
	require.NotEmpty(t, resultUnread)
}

func TestExists(t *testing.T) {
	ctx := context.Background()

	numbers := []string{""}
	result, err := client.Exists(ctx, os.Getenv("INSTANCE_NAME"), numbers)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
