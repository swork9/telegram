package telegram

import (
	"net/url"
	"strconv"
)

func (t *BotT) ForwardMessage(chatID, fromChatID, messageID int64, disableNotification bool) (*MessageT, error) {
	data := url.Values{}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("from_chat_id", strconv.FormatInt(fromChatID, 10))
	data.Add("message_id", strconv.FormatInt(messageID, 10))
	data.Add("disable_notification", strconv.FormatBool(disableNotification))

	return t.sendRawMessage("forwardMessage", data)
}
