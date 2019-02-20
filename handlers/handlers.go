package handlers

import (
	"github.com/gin-gonic/gin"
)

// Tableflip is here.
func Tableflip(c *gin.Context) {
	command := c.PostForm("command")

	if command == "/tableflip" {
		c.JSON(200, gin.H{
			"response_type": "in_channel",
			"text":          "(╯°□°）╯︵ ┻━┻",
		})
	} else {
		c.JSON(200, gin.H{
			"text": command,
		})
	}
}
