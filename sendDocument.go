package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

func (t *BotT) SendDocument(chatID int64, document string, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("document", document)
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("sendDocument", data)
}

func (t *BotT) SendDocumentFromBytes(chatID int64, file []byte, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawFile("sendDocument", data, "document", file)
}

func (t *BotT) SendDocumentFromFile(chatID int64, file string, caption string, keyboard KeyboardI) (*MessageT, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t.SendDocumentFromBytes(chatID, bytes, caption, keyboard)
}

func (u *UpdateT) AnswerSendDocument(chatID int64, document string, caption string, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.SendDocument(chatID, document, caption, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "sendDocument"
	data["document"] = document
	data["caption"] = caption
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}
