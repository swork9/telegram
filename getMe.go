package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type GetMeResultT struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
	Result      *UserT `json:"result"`
}

func (t *BotT) GetMe() (*UserT, error) {
	body, httpStatusCode, err := t.sendRaw("getMe", url.Values{})
	if err != nil {
		return nil, err
	}

	result := &GetMeResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("telegram http status: %d. error: %s", httpStatusCode, err.Error())
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram http status: %d. error code: %d. %s", httpStatusCode, result.ErrorCode, result.Description)
	}

	return result.Result, nil
}
