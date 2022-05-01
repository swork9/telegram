package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type GetChatMemberResultT struct {
	Ok          bool        `json:"ok"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Result      *ChatMember `json:"result"`
}

func (t *BotT) GetChatMember(chatID, userID int64) (*ChatMember, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("user_id", strconv.FormatInt(userID, 10))

	body, httpStatusCode, err := t.sendRaw("getChatMember", data)
	if err != nil {
		return nil, err
	}

	result := &GetChatMemberResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("telegram http status: %d. error: %s", httpStatusCode, err.Error())
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram http status: %d. error code: %d. %s", httpStatusCode, result.ErrorCode, result.Description)
	}

	return result.Result, nil
}
