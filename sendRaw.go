package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

type MethodResultT struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type MessageResultT struct {
	Ok      bool      `json:"ok"`
	Message *MessageT `json:"result"`
}

func (t *BotT) sendRaw(method string, values url.Values) ([]byte, error) {
	resp, err := t.client.PostForm("https://api.telegram.org/bot"+t.Token+"/"+method, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if t.Debug {
		fmt.Println(string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Something wrong with Telegram answer: %d", resp.StatusCode)
	}

	return body, nil
}

func (t *BotT) sendRawMethod(method string, values url.Values) error {
	body, err := t.sendRaw(method, values)
	if err != nil {
		return err
	}

	result := &MethodResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if !result.Ok || !result.Result {
		return fmt.Errorf("Telegram reject our method")
	}

	return nil
}

func (t *BotT) sendRawMessage(method string, values url.Values) (*MessageT, error) {
	body, err := t.sendRaw(method, values)
	if err != nil {
		return nil, err
	}

	result := &MessageResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, fmt.Errorf("Telegram reject our message")
	}

	return result.Message, nil
}

func (t *BotT) sendRawFile(method string, values url.Values, fileid string, file, thumb []byte) (*MessageT, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if file != nil {
		filePart, err := writer.CreateFormFile(fileid, "file")
		if err != nil {
			return nil, err
		}
		filePart.Write(file)
	}

	if thumb != nil {
		thumbPart, err := writer.CreateFormFile("thumb", "attach://file")
		if err != nil {
			return nil, err
		}
		thumbPart.Write(thumb)
	}

	for k, v := range values {
		for _, i := range v {
			writer.WriteField(k, i)
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+t.Token+"/"+method, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if t.Debug {
		fmt.Println(string(bodyContent))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Something wrong with Telegram answer")
	}

	result := &MessageResultT{}
	err = json.Unmarshal(bodyContent, result)
	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, fmt.Errorf("Telegram reject our message")
	}

	return result.Message, nil
}
