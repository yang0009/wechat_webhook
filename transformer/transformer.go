package transformer

import (
	"bytes"
	"fmt"
	"wechat_webhook/model"
)

func TransformToMarkdown(notification model.Notification) (markdown *model.WechatMarkdown, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	// annotations := notification.CommonAnnotations

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	// for _, alert := range notification.Alerts {
	// 	annotations := alert.Annotations
	// 	buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
	// 	buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	// }
	alert := notification.Alerts[0]
	fmt.Println(notification.CommonAnnotations)
	markdown = &model.WechatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: fmt.Sprintf("通知组：%s \n\n>当前状态: %s \n\n >查询语句: %s\n\n > 报警信息: %s\n\n> 开始时间：%s\n", groupKey, status, notification.CommonAnnotations["summary"], notification.CommonAnnotations["description"], alert.StartsAt.Format("15:04:05")),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}
