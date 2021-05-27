package telegram

import "encoding/json"

type MediaGroupItem struct {
	Type    string `json:"type"`
	Media   string `json:"media"`
	Caption string `json:"caption,omitempty"`
}

type MediaGroup []MediaGroupItem

func (m MediaGroup) SetCaption(caption string) {
	if len(m) == 0 {
		return
	}

	m[0].Caption = caption
}

func (m MediaGroup) Get() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "[]"
	}

	return string(bytes)
}

func NewMediaGroup(mediaType string, items []string) MediaGroup {
	mediaGroup := MediaGroup{}
	for _, i := range items {
		mediaGroup = append(mediaGroup, MediaGroupItem{
			Type:  mediaType,
			Media: i,
		})
	}

	return mediaGroup
}
