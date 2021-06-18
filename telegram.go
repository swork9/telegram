package telegram

import (
	"net/http"
)

type MessageHandlerF func(*UpdateT, *MessageT)
type CommandHandlerF func(*UpdateT, *MessageT)
type CallbackHandlerF func(*UpdateT, *UserT, *MessageT, []string)

type BotT struct {
	client *http.Client

	MessageHandlers  []MessageHandlerF
	CommandHandlers  map[string]CommandHandlerF
	CallBackHandlers map[string]CallbackHandlerF

	Token string
	Me    *UserT
	Debug bool
}

func New(token string) (*BotT, error) {
	bot := &BotT{client: &http.Client{}, Token: token}
	bot.MessageHandlers = []MessageHandlerF{}
	bot.CommandHandlers = make(map[string]CommandHandlerF)
	bot.CallBackHandlers = make(map[string]CallbackHandlerF)

	return bot, nil
}
