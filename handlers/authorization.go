package handlers

import (
	"github.com/gin-gonic/gin"
)

// Authorization of Slack OAuth.
func Authorization(c *gin.Context) {
	c.JSON(200, gin.H{
		"response_type": "in_channel",
		"text":          "(╯°□°）╯︵ ┻━┻",
	})
}
