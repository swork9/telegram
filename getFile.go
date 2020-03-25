package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GetFileResultT struct {
	Ok     bool   `json:"ok"`
	Result *FileT `json:"result"`
}

func (t *BotT) GetFile(fileID string) (*FileT, []byte, error) {
	data := url.Values{}
	data.Add("file_id", fileID)

	body, err := t.sendRaw("getFile", data)
	if err != nil {
		return nil, nil, err
	}

	result := &GetFileResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, nil, err
	}

	if !result.Ok {
		return nil, nil, fmt.Errorf("Telegram refuse our getFile request")
	}

	if t.Debug {
		fmt.Println("https://api.telegram.org/file/bot" + t.Token + "/" + result.Result.Path)
	}

	resp, err := t.client.Get("https://api.telegram.org/file/bot" + t.Token + "/" + result.Result.Path)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Something wrong with Telegram file answer: %d", resp.StatusCode)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return result.Result, body, nil
}
