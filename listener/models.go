package listener

import (
	"time"
)

type MessageUpsertListener func(message *MessageUpsert) error

type MessageUpsertDataMessageImageMessage struct {
	Url               string                  `json:"url"`
	Mimetype          string                  `json:"mimetype"`
	FileSha256        string                  `json:"fileSha256"`
	FileLength        string                  `json:"fileLength"`
	Height            int                     `json:"height"`
	Width             int                     `json:"width"`
	MediaKey          string                  `json:"mediaKey"`
	FileEncSha256     string                  `json:"fileEncSha256"`
	DirectPath        string                  `json:"directPath"`
	MediaKeyTimestamp string                  `json:"mediaKeyTimestamp"`
	JpegThumbnail     string                  `json:"jpegThumbnail"`
	ContextInfo       ImageMessageContextInfo `json:"contextInfo"`
	ViewOnce          bool                    `json:"viewOnce"`
}
type ImageMessageContextInfo struct {
	DisappearingMode MessageUpsertDataContextInfoDisappearingMode `json:"disappearingMode"`
}

type MessageUpsertDataMessageMessageContextInfo struct {
	DeviceListMetadata        MessageContextInfoDeviceListMetadata `json:"deviceListMetadata"`
	DeviceListMetadataVersion int                                  `json:"deviceListMetadataVersion"`
	MessageSecret             string                               `json:"messageSecret"`
}

type MessageContextInfoDeviceListMetadata struct {
	SenderKeyHash       string `json:"senderKeyHash"`
	SenderTimestamp     string `json:"senderTimestamp"`
	SenderAccountType   string `json:"senderAccountType"`
	ReceiverAccountType string `json:"receiverAccountType"`
	RecipientKeyHash    string `json:"recipientKeyHash"`
	RecipientTimestamp  string `json:"recipientTimestamp"`
}

type MessageUpsert struct {
	Event       string            `json:"event"`
	Instance    string            `json:"instance"`
	Data        MessageUpsertData `json:"data"`
	Destination string            `json:"destination"`
	DateTime    time.Time         `json:"date_time"`
	Sender      string            `json:"sender"`
	ServerUrl   string            `json:"server_url"`
	Apikey      string            `json:"apikey"`
}

type MessageUpsertData struct {
	Key              MessageUpsertDataKey         `json:"key"`
	PushName         string                       `json:"pushName"`
	Status           string                       `json:"status"`
	Message          MessageUpsertDataMessage     `json:"message"`
	ContextInfo      MessageUpsertDataContextInfo `json:"contextInfo"`
	MessageType      string                       `json:"messageType"`
	MessageTimestamp int                          `json:"messageTimestamp"`
	InstanceId       string                       `json:"instanceId"`
	Source           string                       `json:"source"`
}

type MessageUpsertDataKey struct {
	RemoteJid   string `json:"remoteJid"`
	FromMe      bool   `json:"fromMe"`
	Id          string `json:"id"`
	Participant string `json:"participant"` // Use this field do check if is a group
}

type MessageUpsertDataMessage struct {
	Conversation       string                                     `json:"conversation"`
	Base64             string                                     `json:"base64"`
	ImageMessage       MessageUpsertDataMessageImageMessage       `json:"imageMessage"`
	MessageContextInfo MessageUpsertDataMessageMessageContextInfo `json:"messageContextInfo"`
}

type MessageUpsertDataContextInfo struct {
	EphemeralSettingTimestamp string                                       `json:"ephemeralSettingTimestamp"`
	DisappearingMode          MessageUpsertDataContextInfoDisappearingMode `json:"disappearingMode"`
	StanzaId                  string                                       `json:"stanzaId"`
	Participant               string                                       `json:"participant"`
	Expiration                int                                          `json:"expiration"`
	QuotedMessage             MessageUpsertDataContextInfoQuotedMessage    `json:"quotedMessage"`
	MentionedJid              []string                                     `json:"mentionedJid"`
}

type MessageUpsertDataContextInfoDisappearingMode struct {
	Initiator     string `json:"initiator"`
	Trigger       string `json:"trigger"`
	InitiatedByMe bool   `json:"initiatedByMe"`
}

type MessageUpsertDataContextInfoQuotedMessage struct {
	ExtendedTextMessage struct {
		Text        string `json:"text"`
		ContextInfo struct {
			Expiration       int `json:"expiration"`
			DisappearingMode struct {
				Initiator     string `json:"initiator"`
				Trigger       string `json:"trigger"`
				InitiatedByMe bool   `json:"initiatedByMe"`
			} `json:"disappearingMode"`
		} `json:"contextInfo"`
	} `json:"extendedTextMessage"`
}
