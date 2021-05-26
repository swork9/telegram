package telegram

import (
	"net/url"
	"reflect"
	"strconv"
)

type MessageOptions struct {
	ReplyToMessageID    uint64
	ParseMode           ParseModeT
	DisableNotification bool
	DisableWebPreview   bool
	Keyboard            KeyboardI
}

func (o *MessageOptions) GetMap() map[string]string {
	dataMap := map[string]string{}

	if o.ReplyToMessageID != 0 {
		dataMap["reply_to_message_id"] = strconv.FormatUint(o.ReplyToMessageID, 10)
	}

	if o.ParseMode == ParseModeMarkdown {
		dataMap["parse_mode"] = "Markdown"
	} else if o.ParseMode == ParseModeHTML {
		dataMap["parse_mode"] = "HTML"
	}

	dataMap["disable_notification"] = strconv.FormatBool(o.DisableNotification)
	dataMap["disable_web_page_preview"] = strconv.FormatBool(o.DisableWebPreview)

	if o.Keyboard != nil && !reflect.ValueOf(o.Keyboard).IsNil() {
		dataMap["reply_markup"] = o.Keyboard.Get()
	}

	return dataMap
}

func (o *MessageOptions) Get() url.Values {
	dataMap := o.GetMap()
	data := url.Values{}

	for k, v := range dataMap {
		data.Set(k, v)
	}

	return data
}

func NewMessageOptions() *MessageOptions {
	return &MessageOptions{}
}
