package telegram

import (
	"fmt"
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

func (t *BotT) SendVideoGroup(chatID int64, videos []string, caption string, options *MessageOptions) (*MessageT, error) {
	if len(videos) == 0 {
		return nil, fmt.Errorf("videos slice can't be nil")
	}
	if len(videos) == 1 {
		return t.SendVideo(chatID, videos[0], caption, options)
	}

	data := url.Values{}
	if options != nil {
		data = options.Get()
	}

	data.Add("chat_id", strconv.FormatInt(chatID, 10))

	mediaGroup := NewMediaGroup("video", videos)
	mediaGroup.SetCaption(caption)

	data.Add("media", mediaGroup.Get())

	return t.sendRawMessage("sendMediaGroup", data)
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

func (u *UpdateT) SendVideoGroup(chatID int64, videos []string, caption string, options *MessageOptions) {
	if u.context == nil {
		_, _ = u.bot.SendVideoGroup(chatID, videos, caption, options)
		return
	}

	if len(videos) == 0 {
		return
	}
	if len(videos) == 1 {
		u.SendVideo(chatID, videos[0], caption, options)
		return
	}

	data := map[string]string{}
	if options != nil {
		data = options.GetMap()
	}

	data["method"] = "sendMediaGroup"

	mediaGroup := NewMediaGroup("video", videos)
	mediaGroup.SetCaption(caption)

	data["media"] = mediaGroup.Get()

	u.context.JSON(http.StatusOK, data)
}
