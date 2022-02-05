package talk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/f97one/LineWorksBotTalker/v1/jwt"
	"github.com/f97one/LineWorksBotTalker/v1/settings"
	"net/http"
	"time"
)

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Payload struct {
	AccountId *string `json:"account_id"`
	RoomId    *string `json:"room_id"`
	Content   Content `json:"content"`
}

func NewTextPayload(accountId *string, roomId *string, msg string) Payload {
	return Payload{
		AccountId: accountId,
		RoomId:    roomId,
		Content: Content{
			Type: "text",
			Text: msg,
		},
	}
}

func SendText(accessToken string, config settings.LWBotTalkConfig, payload Payload) error {
	textEndpoint := fmt.Sprintf("https://apis.worksmobile.com/r/%s/message/v1/bot/%d/message/push", config.ApiId, config.BotNo)

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, textEndpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("consumerKey", config.ConsumerKey)

	client := &http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	err = jwt.ParseStateError(resp)
	if err != nil {
		return err
	}

	return nil
}
