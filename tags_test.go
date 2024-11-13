package evolution_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTagsMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.GetTags(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestAddTag(t *testing.T) {
	ctx := context.Background()
	err := client.AddChatTag(ctx, os.Getenv("INSTANCE_NAME"), "7", os.Getenv("NUMBER"))
	require.NoError(t, err)
}
