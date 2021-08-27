package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func (t *BotT) SendAudio(chatID int64, audio string, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("audio", audio)
	data.Add("caption", caption)

	return t.sendRawMessage("sendAudio", data)
}

func (t *BotT) SendAudioFromBytes(chatID int64, filename string, file []byte, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	return t.sendRawFile("sendAudio", data, "audio", filename, file, nil)
}

func (t *BotT) SendAudioFromFile(chatID int64, file string, caption string, options *MessageOptions) (*MessageT, error) {
	var fileBytes []byte

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileBytes, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return t.SendAudioFromBytes(chatID, filepath.Base(file), fileBytes, caption, options)
}

func (u *UpdateT) SendAudio(chatID int64, audio string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendAudio(chatID, audio, caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendAudio"
	data["audio"] = audio
	data["caption"] = caption

	u.context.JSON(http.StatusOK, data)
}
