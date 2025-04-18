package listener

import (
	"time"

	"github.com/verbeux-ai/evolution-client-go"
)

type WookType string

const (
	WookTypeMessageUpsert  WookType = "messages.upsert"
	WookTypePresenceUpdate WookType = "presence.update"
)

type WookIdentifier struct {
	Event WookType `json:"event"`
}

type MessageUpsertListener func(message *MessageUpsert) error

type MessageUpsertDataMessageAudioMessage struct {
	Url               string                 `json:"url"`
	Mimetype          string                 `json:"mimetype"`
	FileSha256        string                 `json:"fileSha256"`
	FileLength        string                 `json:"fileLength"`
	Seconds           int                    `json:"seconds"`
	Ptt               bool                   `json:"ptt"`
	MediaKey          string                 `json:"mediaKey"`
	FileEncSha256     string                 `json:"fileEncSha256"`
	DirectPath        string                 `json:"directPath"`
	MediaKeyTimestamp string                 `json:"mediaKeyTimestamp"`
	ContextInfo       FileMessageContextInfo `json:"contextInfo"`
	Waveform          string                 `json:"waveform"`
	ViewOnce          bool                   `json:"viewOnce"`
}
type MessageUpsertDataMessageImageMessage struct {
	Url               string                 `json:"url"`
	Mimetype          string                 `json:"mimetype"`
	FileSha256        string                 `json:"fileSha256"`
	FileLength        string                 `json:"fileLength"`
	Height            int                    `json:"height"`
	Caption           string                 `json:"caption"`
	Width             int                    `json:"width"`
	MediaKey          string                 `json:"mediaKey"`
	FileEncSha256     string                 `json:"fileEncSha256"`
	DirectPath        string                 `json:"directPath"`
	MediaKeyTimestamp string                 `json:"mediaKeyTimestamp"`
	JpegThumbnail     string                 `json:"jpegThumbnail"`
	ContextInfo       FileMessageContextInfo `json:"contextInfo"`
	ViewOnce          bool                   `json:"viewOnce"`
}
type FileMessageContextInfo struct {
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
	AudioMessage       MessageUpsertDataMessageAudioMessage       `json:"audioMessage"`
	ReactionMessage    ReactionMessage                            `json:"reactionMessage"`
	MessageContextInfo MessageUpsertDataMessageMessageContextInfo `json:"messageContextInfo"`
}

type ReactionMessage struct {
	Key               MessageUpsertDataKey `json:"key"`
	Text              string               `json:"text"`
	SenderTimestampMs string               `json:"senderTimestampMs"`
}

type MessageUpsertDataContextInfo struct {
	EphemeralSettingTimestamp                   string                                       `json:"ephemeralSettingTimestamp"`
	DisappearingMode                            MessageUpsertDataContextInfoDisappearingMode `json:"disappearingMode"`
	StanzaId                                    string                                       `json:"stanzaId"`
	Participant                                 string                                       `json:"participant"`
	Expiration                                  int                                          `json:"expiration"`
	QuotedMessage                               MessageUpsertDataContextInfoQuotedMessage    `json:"quotedMessage"`
	MentionedJid                                []string                                     `json:"mentionedJid"`
	ConversionSource                            string                                       `json:"conversionSource"`
	ConversionData                              string                                       `json:"conversionData"`
	ConversionDelaySeconds                      int                                          `json:"conversionDelaySeconds"`
	MessageUpsertDataContextInfoExternalAdReply MessageUpsertDataContextInfoExternalAdReply  `json:"externalAdReply"`
	EntryPointConversionSource                  string                                       `json:"entryPointConversionSource"`
	EntryPointConversionApp                     string                                       `json:"entryPointConversionApp"`
	EntryPointConversionDelaySeconds            int                                          `json:"entryPointConversionDelaySeconds"`
	TrustBannerAction                           int64                                        `json:"trustBannerAction"`
}

type MessageUpsertDataContextInfoExternalAdReply struct {
	Title                 string `json:"title"`
	Body                  string `json:"body"`
	MediaType             string `json:"mediaType"`
	ThumbnailUrl          string `json:"thumbnailUrl"`
	Thumbnail             string `json:"thumbnail"`
	SourceType            string `json:"sourceType"`
	SourceId              string `json:"sourceId"`
	SourceUrl             string `json:"sourceUrl"`
	ContainsAutoReply     bool   `json:"containsAutoReply"`
	RenderLargerThumbnail bool   `json:"renderLargerThumbnail"`
	ShowAdAttribution     bool   `json:"showAdAttribution"`
	CtwaClid              string `json:"ctwaClid"`
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

type PresenceUpdateListener func(data *PresenceUpdate) error

type PresenceUpdate struct {
	Instance    string             `json:"instance"`
	Data        PresenceUpdateData `json:"data"`
	Destination string             `json:"destination"`
	DateTime    time.Time          `json:"date_time"`
	Sender      string             `json:"sender"`
	ServerUrl   string             `json:"server_url"`
	Apikey      interface{}        `json:"apikey"`
}

type PresenceUpdateData struct {
	Id        string                                             `json:"id"`
	Presences map[string]PresenceUpdateDataPresencesSWhatsappNet `json:"presences"`
}

type PresenceUpdateDataPresencesSWhatsappNet struct {
	LastKnownPresence evolution.SendPresenceRequestPresence `json:"lastKnownPresence"`
}
