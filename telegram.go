package telegram

import (
	"net/http"
	"strconv"
	"strings"
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
	MyID  int64
	Debug bool
}

func parseMyID(token string) int64 {
	s := strings.Split(token, ":")
	if len(s) == 0 {
		return 0
	}

	id, _ := strconv.ParseInt(s[0], 10, 64)

	return id
}

func New(token string) (*BotT, error) {
	bot := &BotT{client: &http.Client{}, Token: token, MyID: parseMyID(token)}
	bot.MessageHandlers = []MessageHandlerF{}
	bot.CommandHandlers = make(map[string]CommandHandlerF)
	bot.CallBackHandlers = make(map[string]CallbackHandlerF)

	return bot, nil
}
