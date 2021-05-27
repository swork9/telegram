package telegram

import "encoding/json"

type MediaGroupItem struct {
	Type    string `json:"type"`
	Media   string `json:"media"`
	Caption string `json:"caption,omitempty"`
}

type MediaGroup struct {
	items []MediaGroupItem
}

func (m *MediaGroup) Len() int {
	return len(m.items)
}

func (m *MediaGroup) Add(mediaType, mediaID string) {
	m.items = append(m.items, MediaGroupItem{Type: mediaType, Media: mediaID})
}

func (m *MediaGroup) SetCaption(caption string) {
	if len(m.items) == 0 {
		return
	}

	m.items[0].Caption = caption
}

func (m *MediaGroup) Get() string {
	bytes, err := json.Marshal(m.items)
	if err != nil {
		return "[]"
	}

	return string(bytes)
}

func NewMediaGroup() *MediaGroup {
	return &MediaGroup{}
}
