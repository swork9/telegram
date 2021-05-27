package telegram

import (
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
