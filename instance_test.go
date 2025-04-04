package evolution_test

import (
	"context"
	"os"
	"testing"

	"github.com/verbeux-ai/evolution-client-go"

	"github.com/stretchr/testify/require"
)

func TestCreateInstance(t *testing.T) {
	ctx := context.Background()

	result, err := client.CreateInstance(ctx, &evolution.CreateInstanceRequest{
		InstanceName:    os.Getenv("INSTANCE_NAME"),
		Qrcode:          false,
		Integration:     "WHATSAPP-BAILEYS",
		RejectCall:      false,
		MsgCall:         "",
		GroupsIgnore:    false,
		AlwaysOnline:    false,
		ReadMessages:    false,
		ReadStatus:      false,
		SyncFullHistory: false,
		ProxyHost:       "",
		ProxyPort:       "",
		ProxyProtocol:   "",
		ProxyUsername:   "",
		ProxyPassword:   "",
		Webhook: &evolution.CreateInstanceRequestWebhook{
			Url:      "https://webhook.site/89930b26-a76e-425b-a59f-36f925f0863c",
			ByEvents: false,
			Base64:   true,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Events: []string{
				"APPLICATION_STARTUP",
				"QRCODE_UPDATED",
				"MESSAGES_SET",
				"MESSAGES_UPSERT",
				"MESSAGES_UPDATE",
				"MESSAGES_DELETE",
				"SEND_MESSAGE",
				"CONTACTS_SET",
				"CONTACTS_UPSERT",
				"CONTACTS_UPDATE",
				"PRESENCE_UPDATE",
				"CHATS_SET",
				"CHATS_UPSERT",
				"CHATS_UPDATE",
				"CHATS_DELETE",
				"GROUPS_UPSERT",
				"GROUP_UPDATE",
				"GROUP_PARTICIPANTS_UPDATE",
				"CONNECTION_UPDATE",
				"LABELS_EDIT",
				"LABELS_ASSOCIATION",
				"CALL",
				"TYPEBOT_START",
				"TYPEBOT_CHANGE_STATUS",
			},
		},
		Rabbitmq:                        &evolution.CreateInstanceRequestRabbitMQ{},
		Sqs:                             &evolution.CreateInstanceRequestSqs{},
		ChatwootAccountId:               "",
		ChatwootToken:                   "",
		ChatwootUrl:                     "",
		ChatwootSignMsg:                 false,
		ChatwootReopenConversation:      false,
		ChatwootConversationPending:     false,
		ChatwootImportContacts:          false,
		ChatwootNameInbox:               "",
		ChatwootMergeBrazilContacts:     false,
		ChatwootImportMessages:          false,
		ChatwootDaysLimitImportMessages: 0,
		ChatwootOrganization:            "",
		ChatwootLogo:                    "",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestRestartInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.RestartInstance(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestLogoutInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.LogoutInstance(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestDeleteInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.DeleteInstance(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestConnectInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.ConnectInstance(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetConnectionStateInstance(t *testing.T) {
	ctx := context.Background()
	result, err := client.ConnectionStateInstance(ctx, os.Getenv("INSTANCE_NAME"))
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestFetchInstances(t *testing.T) {
	ctx := context.Background()
	result, err := client.FetchInstances(ctx, nil)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestFetchInstancesFiltered(t *testing.T) {
	ctx := context.Background()
	result, err := client.FetchInstances(ctx, &evolution.FetchInstancesRequestFilter{
		InstanceName: "5459c0a2-a35f-11ef-bfad-aa807b84a752",
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
