package telegram

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type ParseModeT int

const (
	ParseModeNone ParseModeT = iota
	ParseModeMarkdown
	ParseModeHTML
)

func (t *BotT) SendMessage(chatID int64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("text", text)
	if parseMode == ParseModeMarkdown {
		data.Add("parse_mode", "Markdown")
	} else if parseMode == ParseModeHTML {
		data.Add("parse_mode", "HTML")
	}
	data.Add("disable_web_page_preview", strconv.FormatBool(disableWebPreview))

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("sendMessage", data)
}

func (t *BotT) ReplyMessage(chatID int64, replyMessageID uint64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("reply_to_message_id", strconv.FormatUint(replyMessageID, 10))
	data.Add("text", text)
	if parseMode == ParseModeMarkdown {
		data.Add("parse_mode", "Markdown")
	} else if parseMode == ParseModeHTML {
		data.Add("parse_mode", "HTML")
	}
	data.Add("disable_web_page_preview", strconv.FormatBool(disableWebPreview))

	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("sendMessage", data)
}

func (t *BotT) EditMessageText(chatID int64, messageID uint64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("message_id", strconv.FormatUint(messageID, 10))
	data.Add("text", text)
	if parseMode == ParseModeMarkdown {
		data.Add("parse_mode", "Markdown")
	} else if parseMode == ParseModeHTML {
		data.Add("parse_mode", "HTML")
	}
	data.Add("disable_web_page_preview", strconv.FormatBool(disableWebPreview))
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("editMessageText", data)
}

func (t *BotT) EditMessageCaption(chatID int64, messageID uint64, caption string, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("message_id", strconv.FormatUint(messageID, 10))
	data.Add("caption", caption)
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("editMessageCaption", data)
}

func (t *BotT) EditMessageKeyboard(chatID int64, messageID uint64, keyboard KeyboardI) (*MessageT, error) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("message_id", strconv.FormatUint(messageID, 10))
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data.Add("reply_markup", keyboard.Get())
	}

	return t.sendRawMessage("editMessageReplyMarkup", data)
}

func (t *BotT) DeleteMessage(chatID int64, messageID uint64) {
	data := url.Values{}
	data.Add("chat_id", strconv.FormatInt(chatID, 10))
	data.Add("message_id", strconv.FormatUint(messageID, 10))

	t.sendRawMethod("deleteMessage", data)
}

func (t *BotT) CallbackQuery(callbackID string, text string) {
	data := url.Values{}
	data.Add("callback_query_id", callbackID)
	data.Add("text", text)

	t.sendRawMethod("answerCallbackQuery", data)
}

func (u *UpdateT) AnswerSendMessage(chatID int64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard *InlineKeyboardMarkupT) {
	if u.context == nil {
		u.bot.SendMessage(chatID, text, parseMode, disableWebPreview, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "sendMessage"
	data["chat_id"] = chatID
	data["text"] = text
	if parseMode == ParseModeMarkdown {
		data["parse_mode"] = "Markdown"
	} else if parseMode == ParseModeHTML {
		data["parse_mode"] = "HTML"
	}
	data["disable_web_page_preview"] = disableWebPreview
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerReplyMessage(chatID int64, replyMessageID uint64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard *InlineKeyboardMarkupT) {
	if u.context == nil {
		u.bot.ReplyMessage(chatID, replyMessageID, text, parseMode, disableWebPreview, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "sendMessage"
	data["chat_id"] = chatID
	data["reply_to_message_id"] = replyMessageID
	data["text"] = text
	if parseMode == ParseModeMarkdown {
		data["parse_mode"] = "Markdown"
	} else if parseMode == ParseModeHTML {
		data["parse_mode"] = "HTML"
	}
	data["disable_web_page_preview"] = disableWebPreview
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerEditMessageText(chatID int64, messageID uint64, text string, parseMode ParseModeT, disableWebPreview bool, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.EditMessageText(chatID, messageID, text, parseMode, disableWebPreview, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "editMessageText"
	data["chat_id"] = chatID
	data["message_id"] = messageID
	data["text"] = text
	if parseMode == ParseModeMarkdown {
		data["parse_mode"] = "Markdown"
	} else if parseMode == ParseModeHTML {
		data["parse_mode"] = "HTML"
	}
	data["disable_web_page_preview"] = disableWebPreview
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerEditMessageCaption(chatID int64, messageID uint64, caption string, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.EditMessageCaption(chatID, messageID, caption, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "editMessageCaption"
	data["chat_id"] = chatID
	data["message_id"] = messageID
	data["caption"] = caption
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerEditMessageKeyboard(chatID int64, messageID uint64, keyboard KeyboardI) {
	if u.context == nil {
		u.bot.EditMessageKeyboard(chatID, messageID, keyboard)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "editMessageReplyMarkup"
	data["chat_id"] = chatID
	data["message_id"] = messageID
	if keyboard != nil && !reflect.ValueOf(keyboard).IsNil() {
		data["reply_markup"] = keyboard.Get()
	}

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerDeleteMessage(chatID int64, messageID uint64) {
	if u.context == nil {
		u.bot.DeleteMessage(chatID, messageID)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "deleteMessage"
	data["chat_id"] = chatID
	data["message_id"] = messageID

	u.context.JSON(http.StatusOK, data)
}

func (u *UpdateT) AnswerCallbackQuery(text string) {
	if u.Callback == nil {
		return
	}
	if u.context == nil {
		u.bot.CallbackQuery(u.Callback.ID, text)
		return
	}

	data := map[string]interface{}{}
	data["method"] = "answerCallbackQuery"
	data["callback_query_id"] = u.Callback.ID
	data["text"] = text

	u.context.JSON(http.StatusOK, data)
}
