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

func TestSendListMessage(t *testing.T) {
	ctx := context.Background()
	result, err := client.SendListMessage(ctx, os.Getenv("INSTANCE_NAME"), &evolution.ListMessageRequest{
		Number:      os.Getenv("NUMBER"),
		Title:       "Testando",
		Description: "Descricao",
		ButtonText:  "Botao text",
		FooterText:  "Footer text",
		Sections: []evolution.ListMessageRequestSection{
			{
				Title: "teste title",
				Rows: []evolution.ListMessageRequestSectionRow{
					{
						Title:       "teste row title",
						Description: "teste row description",
						RowID:       "1",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
