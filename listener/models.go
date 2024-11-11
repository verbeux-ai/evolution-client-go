package listener

import (
	"time"
)

type MessageUpsertListener func(message *MessageUpsert) error

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
	Conversation string `json:"conversation"`
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
