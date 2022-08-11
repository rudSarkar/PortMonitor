package slackwebhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SlackSendMessage(webHookUrl string, message string) error {

	jsonBody, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", webHookUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
