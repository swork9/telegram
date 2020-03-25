package telegram

import (
	"encoding/json"
)

type KeyboardI interface {
	Get() string
}

type InlineKeyboardMarkupT struct {
	Buttons [][]*InlineKeyboardButtonT `json:"inline_keyboard"`
}

type InlineKeyboardButtonT struct {
	Text string `json:"text"`
	Url  string `json:"url"`
	Data string `json:"callback_data"`
}

func (k *InlineKeyboardMarkupT) Get() string {
	bytes, err := json.Marshal(k)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}
