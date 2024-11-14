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
			RemoteJID: "558598437440",
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
