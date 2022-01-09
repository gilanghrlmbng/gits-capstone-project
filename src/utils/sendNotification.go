package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"src/config"

	"github.com/labstack/echo/v4"
)

type Notification struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	ClickAction string `json:"click_action,omitempty"`
}
type RequestSendNotificationToken struct {
	To           string       `json:"to"`
	Notification Notification `json:"notification"`
}

func SendNotificationToken(c echo.Context, reqData RequestSendNotificationToken) error {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(reqData)

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", payloadBuf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", config.GetConfigs(c).FirebaseApiKey))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
