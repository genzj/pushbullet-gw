package pushbullet

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
)

type accessTokenRequest struct {
	ClientCredential
	GrantType string `json:"grant_type"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (c *Client) RefreshToken() error {
	if c.Credential.Code == "" {
		log.Errorf("code is required to get token")
		return fmt.Errorf("code is required to get token")
	}

	req := accessTokenRequest{
		ClientCredential: c.Credential,
		GrantType:        "authorization_code",
	}

	if resp, err := resty.R().
		SetBody(req).
		SetResult(&tokenResponse{}).
		Post(c.ActionURL("/oauth2/token").String()); err != nil {
		log.Error(err)
		return err
	} else if resp.StatusCode()/100 != 2 {
		log.Error(resp.String())
		return fmt.Errorf(resp.String())
	} else {
		log.Debug("token received: ", resp)
		c.LoadToken(resp.Result().(*tokenResponse).AccessToken)
	}
	return nil
}

func (c *Client) LoadToken(token string) {
	c.token = token
}

func (c *Client) HasToken() bool {
	return c.token != ""
}
