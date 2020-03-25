package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type GetMeResultT struct {
	Ok     bool   `json:"ok"`
	Result *UserT `json:"result"`
}

func (t *BotT) GetMe() (*UserT, error) {
	body, err := t.sendRaw("getMe", url.Values{})
	if err != nil {
		return nil, err
	}

	result := &GetMeResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, fmt.Errorf("Telegram refuse our getMe request")
	}

	return result.Result, nil
}
