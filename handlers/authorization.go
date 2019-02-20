package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	slackOauthConfig = &oauth2.Config{
		RedirectURL:  "https://slack.com/oauth/authorize",
		ClientID:     os.Getenv("SLACK_CLIENT_ID"),
		ClientSecret: os.Getenv("SLACK_SECRET"),
		Scopes:       []string{"chat:write:user"},
	}

	secret = os.Getenv("SLACK_SIGNING_SECRET")
)

// Authorization of Slack OAuth.
func Authorization(c *gin.Context) {
	url := slackOauthConfig.AuthCodeURL(secret)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Callback for token.
func Callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	message := "Token: " + code + " State: " + state
	c.String(http.StatusOK, message)
}
