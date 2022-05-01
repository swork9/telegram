package telegram

import "github.com/gin-gonic/gin"

type ChatMemberStatus string

const (
	ChatMemberOwner         ChatMemberStatus = "creator"
	ChatMemberAdministrator                  = "administrator"
	ChatMemberMember                         = "member"
	ChatMemberRestricted                     = "restricted"
	ChatMemberLeft                           = "left"
	ChatMemberBanned                         = "kicked"
)

type UpdateT struct {
	bot     *BotT
	context *gin.Context

	UpdateID uint64     `json:"update_id"`
	Message  *MessageT  `json:"message"`
	Callback *CallbackT `json:"callback_query"`
}

type UserT struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	Firstname    string `json:"first_name"`
	Lastname     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type ChatMember struct {
	User   *UserT           `json:"user"`
	Status ChatMemberStatus `json:"status"`
}

type ChatT struct {
	ID        int64  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type ForwardFromT struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
}

type MessageT struct {
	ID           uint64        `json:"message_id"`
	Date         uint64        `json:"date"`
	From         *UserT        `json:"from"`
	ForwardFrom  *ForwardFromT `json:"forward_from"`
	Chat         *ChatT        `json:"chat"`
	Photo        []*PhotoT     `json:"photo"`
	Document     *DocumentT    `json:"document"`
	Video        *VideoT       `json:"video"`
	Sticker      *StickerT     `json:"sticker"`
	Entities     []*EntityT    `json:"entities"`
	MediaGroupID string        `json:"media_group_id"`
	Text         string        `json:"text"`
	Caption      string        `json:"caption"`
}

type FileT struct {
	ID   string `json:"file_id"`
	Size uint64 `json:"file_size"`
	Path string `json:"file_path"`
}

type PhotoT struct {
	ID     string `json:"file_id"`
	Size   uint64 `json:"file_size"`
	Width  uint64 `json:"width"`
	Height uint64 `json:"height"`
}

type DocumentT struct {
	ID       string  `json:"file_id"`
	Thumb    *PhotoT `json:"thumb"`
	Name     string  `json:"file_name"`
	MimeType string  `json:"mime_type"`
	Size     uint64  `json:"file_size"`
}

type VideoT struct {
	ID       string  `json:"file_id"`
	Width    uint64  `json:"width"`
	Height   uint64  `json:"height"`
	Duration uint64  `json:"duration"`
	Thumb    *PhotoT `json:"thumb"`
	MimeType string  `json:"mime_type"`
	Size     uint64  `json:"file_size"`
}

type StickerT struct {
	ID       string  `json:"file_id"`
	Emoji    string  `json:"emoji"`
	PackName string  `json:"set_name"`
	Width    uint64  `json:"width"`
	Height   uint64  `json:"height"`
	Thumb    *PhotoT `json:"thumb"`
	Size     uint64  `json:"file_size"`
}

type CallbackT struct {
	ID           string    `json:"id"`
	ChatInstance string    `json:"chat_instance"`
	From         *UserT    `json:"from"`
	Message      *MessageT `json:"message"`
	Data         string    `json:"data"`
}

type EntityT struct {
	Offset uint   `json:"offset"`
	Length uint   `json:"length"`
	Type   string `json:"type"`
}
