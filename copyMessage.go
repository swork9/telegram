package telegram

import (
	"net/url"
	"strconv"
)

func (t *BotT) CopyMessage(chatID, fromChatID int64, messageID uint64, newCaption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("from_chat_id", strconv.FormatInt(fromChatID, 10))
	data.Add("message_id", strconv.FormatUint(messageID, 10))

	if len(newCaption) > 0 {
		data.Add("caption", newCaption)
	}

	return t.sendRawMessage("copyMessage", data)
}
