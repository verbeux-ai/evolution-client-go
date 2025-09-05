package listener

import (
	"time"

	"github.com/verbeux-ai/evolution-client-go"
)

type WookType string

const (
	WookTypeMessageUpsert  WookType = "messages.upsert"
	WookTypePresenceUpdate WookType = "presence.update"
	WookTypeMessageUpdate  WookType = "messages.update"
	WookTypeContactUpdate  WookType = "contacts.update"
	WookTypeContactUpsert  WookType = "contacts.upsert"
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
	RemoteLid   string `json:"remoteLid"`
	FromMe      bool   `json:"fromMe"`
	Id          string `json:"id"`
	Participant string `json:"participant"`
}

type MessageUpsertDataMessage struct {
	Conversation        string                                     `json:"conversation"`
	Base64              string                                     `json:"base64"`
	MediaURL            string                                     `json:"mediaUrl"`
	ImageMessage        MessageUpsertDataMessageImageMessage       `json:"imageMessage"`
	DocumentMessage     MessageUpsertDataMessageDocumentMessage    `json:"documentMessage"`
	VideoMessage        MessageUpsertDataMessageVideoMessage       `json:"videoMessage"`
	AudioMessage        MessageUpsertDataMessageAudioMessage       `json:"audioMessage"`
	ReactionMessage     ReactionMessage                            `json:"reactionMessage"`
	MessageContextInfo  MessageUpsertDataMessageMessageContextInfo `json:"messageContextInfo"`
	ListResponseMessage MessageUpsertDataMessageListMessage        `json:"listResponseMessage"`
	ExtendedMessage     *MessageUpsertExtendedTextMessage          `json:"extendedMessage"`
}

type MessageUpsertDataMessageListMessage struct {
	Title             string                                           `json:"title"`
	ListType          string                                           `json:"listType"`
	SingleSelectReply MessageUpsertDataMessageListSingleSelectReply    `json:"singleSelectReply"`
	ContextInfo       MessageUpsertDataMessageListContextInfo          `json:"contextInfo"`
	Description       string                                           `json:"description"`
	ButtonText        string                                           `json:"buttonText"`
	Sections          []MessageUpsertDataMessageListMessageListSection `json:"sections"`
	FooterText        string                                           `json:"footerText"`
}

type MessageUpsertDataMessageListSingleSelectReply struct {
	SelectedRowId string `json:"selectedRowId"`
}

type MessageUpsertDataMessageListContextInfo struct {
	StanzaId      string                                               `json:"stanzaId"`
	Participant   string                                               `json:"participant"`
	QuotedMessage MessageUpsertDataMessageListContextInfoQuotedMessage `json:"quotedMessage"`
}

type MessageUpsertDataMessageListContextInfoQuotedMessage struct {
	MessageContextInfo struct{}                                      `json:"messageContextInfo"`
	ListMessage        MessageUpsertDataMessageListQuotedMessageList `json:"listMessage"`
}

type MessageUpsertDataMessageListQuotedMessageList struct {
	Title       string                                      `json:"title"`
	Description string                                      `json:"description"`
	ButtonText  string                                      `json:"buttonText"`
	ListType    string                                      `json:"listType"`
	Sections    []MessageUpsertDataMessageListQuotedSection `json:"sections"`
	FooterText  string                                      `json:"footerText"`
}

type MessageUpsertDataMessageListQuotedSection struct {
	Title string                                         `json:"title"`
	Rows  []MessageUpsertDataMessageListQuotedSectionRow `json:"rows"`
}

type MessageUpsertDataMessageListQuotedSectionRow struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	RowId       string `json:"rowId"`
}

type MessageUpsertDataMessageVideoMessage struct {
	Url           string `json:"url,omitempty"`
	Mimetype      string `json:"mimetype,omitempty"`
	Caption       string `json:"caption,omitempty"`
	FileSha256    string `json:"fileSha256,omitempty"`
	FileLength    string `json:"fileLength,omitempty"`
	Seconds       uint32 `json:"seconds,omitempty"`
	MediaKey      string `json:"mediaKey,omitempty"`
	FileEncSha256 string `json:"fileEncSha256,omitempty"`
	JPEGThumbnail string `json:"jpegThumbnail,omitempty"`
	GIFPlayback   bool   `json:"gifPlayback,omitempty"`
}

type MessageUpsertDataMessageDocumentMessage struct {
	Url               string `json:"url"`
	Mimetype          string `json:"mimetype"`
	Title             string `json:"title"`
	FileSha256        string `json:"fileSha256"`
	FileLength        string `json:"fileLength"`
	PageCount         int    `json:"pageCount"`
	MediaKey          string `json:"mediaKey"`
	FileName          string `json:"fileName"`
	FileEncSha256     string `json:"fileEncSha256"`
	DirectPath        string `json:"directPath"`
	MediaKeyTimestamp string `json:"mediaKeyTimestamp"`
	ContactVcard      bool   `json:"contactVcard"`
	JpegThumbnail     string `json:"jpegThumbnail"`
	Caption           string `json:"caption"`
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
	QuotedMessage                               MessageUpsertDataMessage                     `json:"quotedMessage"`
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

type MessageUpsertExtendedTextMessage struct {
	Text        string                       `json:"text"`
	ContextInfo MessageUpsertDataContextInfo `json:"contextInfo"`
}

type MessageUpsertDataContextInfoQuotedMessageList struct {
	Title       string                                           `json:"title"`
	Description string                                           `json:"description"`
	ButtonText  string                                           `json:"buttonText"`
	ListType    string                                           `json:"listType"`
	Sections    []MessageUpsertDataMessageListMessageListSection `json:"sections"`
	FooterText  string                                           `json:"footerText"`
}
type MessageUpsertDataMessageListMessageListSection struct {
	Title string                                              `json:"title"`
	Rows  []MessageUpsertDataMessageListMessageListSectionRow `json:"rows"`
}

type MessageUpsertDataMessageListMessageListSectionRow struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	RowId       string `json:"rowId"`
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

type MessageUpdateListener func(data *MessageUpdate) error

type MessageUpdate struct {
	Event       string            `json:"event"`
	Instance    string            `json:"instance"`
	Data        MessageUpdateData `json:"data"`
	Destination string            `json:"destination"`
	DateTime    time.Time         `json:"date_time"`
	Sender      string            `json:"sender"`
	ServerUrl   string            `json:"server_url"`
}

type MessageUpdateDataStatus string

const (
	MessageStatusDeliveryAck MessageUpdateDataStatus = "DELIVERY_ACK"
	MessageStatusRead        MessageUpdateDataStatus = "READ"
)

type MessageUpdateData struct {
	MessageId   string                  `json:"messageId"`
	KeyId       string                  `json:"keyId"`
	RemoteJid   string                  `json:"remoteJid"`
	RemoteLid   string                  `json:"remoteLid"`
	FromMe      bool                    `json:"fromMe"`
	Participant string                  `json:"participant"`
	Status      MessageUpdateDataStatus `json:"status"`
	InstanceId  string                  `json:"instanceId"`
}

type Contact struct {
	RemoteJid     string `json:"remoteJid"`
	RemoteLid     string `json:"remoteLid"`
	PushName      string `json:"pushName"`
	ProfilePicUrl string `json:"profilePicUrl"`
	InstanceId    string `json:"instanceId"`
}

type ContactUpdateUpsert struct {
	Event       string    `json:"event"`
	Instance    string    `json:"instance"`
	Data        []Contact `json:"data"`
	Destination string    `json:"destination"`
	DateTime    time.Time `json:"date_time"`
	Sender      string    `json:"sender"`
	ServerUrl   string    `json:"server_url"`
}

type ContactUpdateListener func(contacts []Contact) error
type ContactUpsertListener func(contacts []Contact) error
