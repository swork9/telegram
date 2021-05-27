package telegram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func (t *BotT) SendPhoto(chatID int64, photo string, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("photo", photo)
	data.Add("caption", caption)

	return t.sendRawMessage("sendPhoto", data)
}

func (t *BotT) SendPhotoFromBytes(chatID int64, filename string, file []byte, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	return t.sendRawFile("sendPhoto", data, "photo", filename, file, nil)
}

func (t *BotT) SendPhotoFromFile(chatID int64, filename string, file string, caption string, options *MessageOptions) (*MessageT, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t.SendPhotoFromBytes(chatID, filename, bytes, caption, options)
}

func (t *BotT) SendPhotoGroup(chatID int64, photos []string, caption string, options *MessageOptions) ([]*MessageT, error) {
	if len(photos) == 0 {
		return nil, fmt.Errorf("photos slice can't be nil")
	}
	if len(photos) == 1 {
		r, err := t.SendPhoto(chatID, photos[0], caption, options)
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

	mediaGroup := NewMediaGroup("photo", photos)
	mediaGroup.SetCaption(caption)

	data.Add("media", mediaGroup.Get())

	return t.sendRawMessageGroup("sendMediaGroup", data)
}

func (u *UpdateT) SendPhoto(chatID int64, photo string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendPhoto(chatID, photo, caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendPhoto"
	data["photo"] = photo
	data["caption"] = caption

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) SendPhotoGroup(chatID int64, photos []string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendPhotoGroup(chatID, photos, caption, options)
		return
	}

	if len(photos) == 0 {
		return
	}
	if len(photos) == 1 {
		u.SendPhoto(chatID, photos[0], caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendMediaGroup"

	mediaGroup := NewMediaGroup("photo", photos)
	mediaGroup.SetCaption(caption)

	data["media"] = mediaGroup.Get()

	u.context.JSON(http.StatusOK, data)
}
