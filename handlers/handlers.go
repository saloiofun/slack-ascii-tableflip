package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Tableflip is here.
func Tableflip(c *gin.Context) {
	token := c.PostForm("token")
	channel := c.PostForm("channel_id")
	methodURL := "https://slack.com/api/chat.postMessage"

	v := url.Values{}
	v.Set("token", token)
	v.Add("channel", channel)
	v.Add("text", "(╯°□°）╯︵ ┻━┻")
	v.Add("as_user", "true")

	resp, err := http.PostForm(methodURL, v)

	if err != nil {
		log.Println("error happened getting the response", err)
		c.String(http.StatusBadRequest, "error happened getting the response", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("error happened reading the body", err)
		return
	}

	log.Println(string(body[:]))

	// command := c.PostForm("command")
	// if command == "/tableflip" {
	// 	c.JSON(200, gin.H{
	// 		"response_type": "in_channel",
	// 		"text":          "(╯°□°）╯︵ ┻━┻",
	// 	})
	// } else {
	// 	c.JSON(200, gin.H{
	// 		"text": command,
	// 	})
	// }
}
