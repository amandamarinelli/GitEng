package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type (
	GitHubClient struct {
		cfg *oauth2.Config
	}
)

func NewGitHubClient(clientId, clientSecret string) *GitHubClient {
	cfg := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/github/callback",
		Endpoint:     github.Endpoint,
	}

	return &GitHubClient{cfg: cfg}
}

func (g *GitHubClient) GithubLogin(c *gin.Context) {
	url := g.cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (g *GitHubClient) GithubCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := g.cfg.Exchange(c, code)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token.AccessToken,
	})
}
