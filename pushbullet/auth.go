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
		Post(c.ActionURL("/oauth2/token").String()); err != nil {
		log.Error(err)
		return err
	} else if resp.StatusCode()/100 != 2 {
		log.Error(resp.String())
		return fmt.Errorf(resp.String())
	} else {
		log.Info(resp)
	}
	return nil
}

func (c *Client) LoadToken(token string) {
	c.token = token
}

func (c *Client) HasToken() bool {
	return c.token != ""
}
