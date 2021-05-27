package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (t *BotT) SendMediaGroup(chatID int64, mediaGroup *MediaGroup, options *MessageOptions) ([]*MessageT, error) {
	if len(mediaGroup.items) == 0 {
		return nil, fmt.Errorf("mediaGroup slice can't be nil")
	}
	if len(mediaGroup.items) == 1 {
		var r *MessageT
		var err error

		switch mediaGroup.items[0].Type {
		case "audio":
			return nil, fmt.Errorf("audio media not yet implemented")
		case "document":
			r, err = t.SendDocument(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		case "photo":
			r, err = t.SendPhoto(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		case "video":
			r, err = t.SendVideo(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		default:
			return nil, fmt.Errorf("unknown media type")
		}

		if err != nil {
			return nil, err
		}

		return []*MessageT{r}, nil
	}

	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("media", mediaGroup.Get())

	return t.sendRawMessageGroup("sendMediaGroup", data)
}

func (u *UpdateT) SendVideoGroup(chatID int64, mediaGroup *MediaGroup, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendMediaGroup(chatID, mediaGroup, options)
		return
	}

	if len(mediaGroup.items) == 0 {
		return
	}
	if len(mediaGroup.items) == 1 {
		switch mediaGroup.items[0].Type {
		case "audio":
			// Not implemented
			return
		case "document":
			u.SendDocument(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		case "photo":
			u.SendPhoto(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		case "video":
			u.SendVideo(chatID, mediaGroup.items[0].Media, mediaGroup.items[0].Caption, options)
		}

		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendMediaGroup"
	data["media"] = mediaGroup.Get()

	u.context.JSON(http.StatusOK, data)
}
