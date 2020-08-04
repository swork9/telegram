package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

func (t *BotT) SendPhoto(chatID int64, photo string, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("photo", photo)
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("sendPhoto", data)
}

func (t *BotT) SendPhotoFromBytes(chatID int64, filename string, file []byte, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawFile("sendPhoto", data, "photo", filename, file, nil)
}

func (t *BotT) SendPhotoFromFile(chatID int64, filename string, file string, caption string, keyboard KeyboardI) (*MessageT, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t.SendPhotoFromBytes(chatID, filename, bytes, caption, keyboard)
}

func (u *UpdateT) AnswerSendPhoto(chatID int64, photo string, caption string, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.SendPhoto(chatID, photo, caption, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "sendPhoto"
	data["photo"] = photo
	data["caption"] = caption
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}
