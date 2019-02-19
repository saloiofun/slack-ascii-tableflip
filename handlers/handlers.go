package handlers

import (
	"github.com/gin-gonic/gin"
)

// Tableflip is here.
func Tableflip(c *gin.Context) {
	if command := c.PostForm("command"); command == "/tableflip" {
		c.JSON(200, gin.H{
			"text": "(╯°□°）╯︵ ┻━┻",
		})
	} else {
		c.JSON(200, gin.H{
			"text": command,
		})
	}
}
