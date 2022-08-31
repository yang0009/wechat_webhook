package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"wechat_webhook/model"
	"wechat_webhook/transformer"
)

func Send(notification model.Notification, wechatRobot string) (err error) {

	markdown, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST",
		wechatRobot,
		bytes.NewBuffer(data))

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}