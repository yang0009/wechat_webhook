package main

import (
	"flag"
	"net/http"
	"wechat_webhook/model"
	"wechat_webhook/notifier"

	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global wechat robot webhook")
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification
		err := c.BindJSON(&notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to wechat successful!"})

	})
	router.Run()
}
