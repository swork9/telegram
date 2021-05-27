package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func (t *BotT) SendDocument(chatID int64, document string, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("document", document)
	data.Add("caption", caption)

	return t.sendRawMessage("sendDocument", data)
}

func (t *BotT) SendDocumentFromBytes(chatID int64, filename string, file []byte, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	return t.sendRawFile("sendDocument", data, "document", filename, file, nil)
}

func (t *BotT) SendDocumentFromFile(chatID int64, file string, caption string, options *MessageOptions) (*MessageT, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t.SendDocumentFromBytes(chatID, filepath.Base(file), bytes, caption, options)
}

func (u *UpdateT) SendDocument(chatID int64, document string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendDocument(chatID, document, caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendDocument"
	data["document"] = document
	data["caption"] = caption

	u.context.JSON(http.StatusOK, data)
}
