package telegram

import (
	"strings"
)

func (t *BotT) proceedUpdate(update *UpdateT) {
	if update.Message != nil {
		if t.MyID == update.Message.From.ID {
			return
		}

		isCommand := false
		if len(update.Message.Entities) > 0 {
			e := update.Message.Entities[0]
			if e.Type == "bot_command" {
				isCommand = true
				command := update.Message.Text[e.Offset+1 : e.Offset+e.Length]
				if handler, ok := t.CommandHandlers[command]; ok {
					handler(update, update.Message)
				}
			}
		}

		if !isCommand {
			for _, handler := range t.MessageHandlers {
				handler(update, update.Message)
			}
		}
	} else if update.Callback != nil {
		data := strings.Split(update.Callback.Data, "|")
		if len(data) > 0 {
			if handler, ok := t.CallBackHandlers[data[0]]; ok {
				handler(update, update.Callback.From, update.Callback.Message, data)
			}
		}
	}
}
