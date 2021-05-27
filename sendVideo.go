package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func (t *BotT) SendVideo(chatID int64, video string, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("video", video)
	data.Add("caption", caption)

	return t.sendRawMessage("sendVideo", data)
}

func (t *BotT) SendVideoFromBytes(chatID int64, filename string, file, thumb []byte, caption string, options *MessageOptions) (*MessageT, error) {
	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	return t.sendRawFile("sendVideo", data, "video", filename, file, thumb)
}

func (t *BotT) SendVideoFromFile(chatID int64, file, thumb string, caption string, options *MessageOptions) (*MessageT, error) {
	var fileBytes, thumbBytes []byte

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileBytes, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if thumb != "" {
		ft, err := os.Open(thumb)
		if err != nil {
			return nil, err
		}
		defer ft.Close()

		thumbBytes, err = ioutil.ReadAll(ft)
		if err != nil {
			return nil, err
		}
	}

	return t.SendVideoFromBytes(chatID, filepath.Base(file), fileBytes, thumbBytes, caption, options)
}

func (u *UpdateT) SendVideo(chatID int64, video string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendVideo(chatID, video, caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendVideo"
	data["video"] = video
	data["caption"] = caption

	u.context.JSON(http.StatusOK, data)
}
