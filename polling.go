package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PollingResultT struct {
	Ok     bool       `json:"ok"`
	Result []*UpdateT `json:"result"`
}

func pollingUpdates(t *BotT, offset uint64, timeout int) (uint64, error) {
	resp, err := t.client.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=%d&offset=%d", t.Token, timeout, offset))
	if err != nil {
		return offset, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return offset, fmt.Errorf("Something wrong with Telegram answer")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return offset, err
	}

	if t.Debug {
		fmt.Println(string(body))
	}

	result := &PollingResultT{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return offset, err
	}

	for _, update := range result.Result {
		update.bot = t
		update.context = nil

		if update.UpdateID >= offset {
			offset = update.UpdateID + 1
		}

		if t.Debug {
			fmt.Printf("------------------------------------------------------\n%s\n%+v\n------------------------------------------------------\n", string(body), update)
		}

		go t.proceedUpdate(update)
	}

	return offset, nil
}

func (t *BotT) StartPolling(timeout int) error {
	var offset uint64
	var err error

	if _, err = t.sendRaw("deleteWebhook", url.Values{}); err != nil {
		return err
	}

	for {
		offset, err = pollingUpdates(t, offset, timeout)
		if err != nil {
			//fmt.Println("Error. Fuck it:", err, reflect.TypeOf(err))
			//time.Sleep(1 * time.Second)
			return err
		}
	}
}
