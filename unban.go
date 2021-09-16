package telegram

import (
	"net/url"
	"strconv"
)

func (t *BotT) UnbanChatMember(chatID, userID int64, onlyIfBanned bool) (bool, error) {
	data := url.Values{}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("user_id", strconv.FormatInt(userID, 10))
	data.Add("only_if_banned", strconv.FormatBool(onlyIfBanned))

	err := t.sendRawMethod("unbanChatMember", data)
	if err != nil {
		return false, err
	}

	return true, err
}
