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
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type MessageResultT struct {
	Ok          bool      `json:"ok"`
	ErrorCode   int       `json:"error_code,omitempty"`
	Description string    `json:"description,omitempty"`
	Message     *MessageT `json:"result,omitempty"`
}

type MessageGroupResultT struct {
	Ok          bool        `json:"ok"`
	ErrorCode   int         `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Message     []*MessageT `json:"result,omitempty"`
}

func (t *BotT) sendRaw(method string, values url.Values) ([]byte, int, error) {
	resp, err := t.client.PostForm("https://api.telegram.org/bot"+t.Token+"/"+method, values)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	if t.Debug {
		fmt.Println(string(body))
	}

	return body, resp.StatusCode, nil
}

func (t *BotT) sendRawMethod(method string, values url.Values) error {
	body, httpStatusCode, err := t.sendRaw(method, values)
	if err != nil {
		return err
	}

	result := &MethodResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("telegram http status: %d. error: %s", httpStatusCode, err.Error())
	}

	if !result.Ok {
		return fmt.Errorf("telegram http status: %d. error code: %d. %s", httpStatusCode, result.ErrorCode, result.Description)
	}

	return nil
}

func (t *BotT) sendRawMessage(method string, values url.Values) (*MessageT, error) {
	body, httpStatusCode, err := t.sendRaw(method, values)
	if err != nil {
		return nil, err
	}

	result := &MessageResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("telegram http status: %d. error: %s", httpStatusCode, err.Error())
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram http status: %d. error code: %d. %s", httpStatusCode, result.ErrorCode, result.Description)
	}

	return result.Message, nil
}

func (t *BotT) sendRawMessageGroup(method string, values url.Values) ([]*MessageT, error) {
	body, httpStatusCode, err := t.sendRaw(method, values)
	if err != nil {
		return nil, err
	}

	result := &MessageGroupResultT{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("telegram http status: %d. error: %s", httpStatusCode, err.Error())
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram http status: %d. error code: %d. %s", httpStatusCode, result.ErrorCode, result.Description)
	}

	return result.Message, nil
}

func (t *BotT) sendRawFile(method string, values url.Values, fileid, filename string, file, thumb []byte) (*MessageT, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if file != nil {
		filePart, err := writer.CreateFormFile(fileid, filename)
		if err != nil {
			return nil, err
		}
		filePart.Write(file)
	}

	if thumb != nil {
		thumbPart, err := writer.CreateFormFile("thumb", "attach://"+filename)
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

	result := &MessageResultT{}
	err = json.Unmarshal(bodyContent, result)
	if err != nil {
		return nil, fmt.Errorf("telegram http status: %d. error: %s", resp.StatusCode, err.Error())
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram http status: %d. error code: %d. %s", resp.StatusCode, result.ErrorCode, result.Description)
	}

	return result.Message, nil
}
