package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (t *BotT) EnableWebhook(webhookURL string) error {
	resp, err := t.client.PostForm("https://api.telegram.org/bot"+t.Token+"/setWebhook", url.Values{"url": {webhookURL}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Something wrong with Telegram answer")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if t.Debug {
		fmt.Println(string(body))
	}

	return nil
}

func (t *BotT) DisableWebhook() error {
	resp, err := t.client.Get("https://api.telegram.org/bot" + t.Token + "/deleteWebhook")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Something wrong with Telegram answer")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if t.Debug {
		fmt.Println(string(body))
	}

	return nil
}

func (t *BotT) Webhook(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		if t.Debug {
			fmt.Println(err)
		}

		return
	}

	update := &UpdateT{}
	err = json.Unmarshal(body, &update)
	if err != nil {
		if t.Debug {
			fmt.Println(err)
		}

		return
	}

	update.bot = t
	update.context = c

	t.proceedUpdate(update)
}
