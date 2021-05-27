package telegram

import (
	"encoding/json"
)

type KeyboardI interface {
	Get() string
}

type KeyboardButton struct {
	Text string `json:"text,omitempty"`
	Url  string `json:"url,omitempty"`
	Data string `json:"callback_data,omitempty"`
}

type KeyboardRow struct {
	buttons []*KeyboardButton
}

func (r *KeyboardRow) AddButton(text, data string) *KeyboardRow {
	r.buttons = append(r.buttons,
		&KeyboardButton{
			Text: text,
			Data: data,
		},
	)

	return r
}

func (r *KeyboardRow) AddURL(text, url string) *KeyboardRow {
	r.buttons = append(r.buttons,
		&KeyboardButton{
			Text: text,
			Url:  url,
		},
	)

	return r
}

func (r *KeyboardRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.buttons)
}

type InlineKeyboard struct {
	Rows []*KeyboardRow `json:"inline_keyboard"`
}

func (k *InlineKeyboard) NewRow() *KeyboardRow {
	row := &KeyboardRow{
		buttons: []*KeyboardButton{},
	}
	k.Rows = append(k.Rows, row)

	return row
}

func (k *InlineKeyboard) Get() string {
	bytes, err := json.Marshal(k)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

func NewInlineKeyboard() *InlineKeyboard {
	return &InlineKeyboard{
		Rows: []*KeyboardRow{},
	}
}
