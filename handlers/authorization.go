package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/slack"
)

var (
	state        = os.Getenv("SLACK_SIGNING_SECRET")
	clientID     = os.Getenv("SLACK_CLIENT_ID")
	clientSecret = os.Getenv("SLACK_SECRET")

	slackOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     slack.Endpoint,
		Scopes:       []string{"chat:write:user"},
		RedirectURL:  "https://guarded-wildwood-89242.herokuapp.com/slack/callback",
	}
)

// Authorization of Slack OAuth.
// https://www.youtube.com/watch?v=OdyXIi6DGYw
// https://medium.com/@pliutau/getting-started-with-oauth2-in-go-2c9fae55d187
func Authorization(c *gin.Context) {
	url := slackOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Callback for token.
// https://golang.org/pkg/net/url/#Values
// https://github.com/solderjs/http-examples/blob/master/go/post-form.go
func Callback(c *gin.Context) {
	code := c.Query("code")
	oauthAccessURL := "https://slack.com/api/oauth.access?"

	v := url.Values{}
	v.Set("client_id", clientID)
	v.Add("client_secret", clientSecret)
	v.Add("code", code)

	resp, err := http.PostForm(oauthAccessURL, v)

	if err != nil {
		log.Println("error happened getting the response", err)
		c.String(http.StatusBadRequest, "error happened getting the response", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("error happened reading the body", err)
		c.String(http.StatusBadRequest, "error happened reading the body", err)
		return
	}

	log.Println(string(body[:]))
	c.String(http.StatusOK, string(body[:]))
}
