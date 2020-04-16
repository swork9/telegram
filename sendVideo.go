package telegram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

func (t *BotT) SendVideo(chatID int64, video string, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("video", video)
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("sendVideo", data)
}

func (t *BotT) SendVideoFromBytes(chatID int64, file, thumb []byte, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("caption", caption)

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawFile("sendVideo", data, "video", file, thumb)
}

func (t *BotT) SendVideoFromFile(chatID int64, file, thumb string, caption string, keyboard KeyboardI) (*MessageT, error) {
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

	return t.SendVideoFromBytes(chatID, fileBytes, thumbBytes, caption, keyboard)
}

func (u *UpdateT) AnswerSendVideo(chatID int64, video string, caption string, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.SendVideo(chatID, video, caption, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "sendVideo"
	data["video"] = video
	data["caption"] = caption
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}
