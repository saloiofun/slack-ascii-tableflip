package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Tableflip is here.
func Tableflip(c *gin.Context) {
	text := strings.Split(c.PostForm("text"), " ")[0]
	ascii := "(╯°□°）╯︵ ┻━┻"

	if text == "down" {
		ascii = "┬─┬ノ( º _ ºノ)"
	}

	c.JSON(200, gin.H{
		"response_type": "in_channel",
		"text":          ascii,
	})

	// channel := c.PostForm("channel_id")
	// text := strings.Split(c.PostForm("text"), " ")[0]
	// methodURL := "https://slack.com/api/chat.postMessage"
	// ascii := "(╯°□°）╯︵ ┻━┻"

	// if text == "down" {
	// 	ascii = "┬─┬ノ( º _ ºノ)"
	// }

	// v := url.Values{}
	// v.Set("token", os.Getenv("SLACK_TOKEN"))
	// v.Add("channel", channel)
	// v.Add("text", ascii)
	// v.Add("as_user", "true")

	// resp, err := http.PostForm(methodURL, v)

	// if err != nil {
	// 	log.Println("error happened getting the response", err)
	// 	c.String(http.StatusBadRequest, "error happened getting the response", err)
	// 	return
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Println("error happened reading the body", err)
	// 	return
	// }

	// log.Println(string(body[:]))
}
