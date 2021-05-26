package telegram

import (
	"net/url"
	"strconv"
)

func (t *BotT) CopyMessage(chatID, fromChatID, messageID int64, disableNotification bool, newCaption string) (*MessageT, error) {
	data := url.Values{}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("from_chat_id", strconv.FormatInt(fromChatID, 10))
	data.Add("message_id", strconv.FormatInt(messageID, 10))
	data.Add("disable_notification", strconv.FormatBool(disableNotification))

	if len(newCaption) > 0 {
		data.Add("caption", newCaption)
	}

	return t.sendRawMessage("copyMessage", data)
}
