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

// LoadToken insert token string to client for next api request
// If the token is in persistency format (as returned by
// TokenToSave), it will be decrypted before being loaded.
func (c *Client) LoadToken(token string) {
	c.token = token
}

func (c *Client) HasToken() bool {
	return c.token != ""
}

// TokenToSave returns a crypted token for persistent storage
func (c *Client) TokenToSave() string {
	return c.token
}

// SafeClone returns a new client with same client id, secret and redirect url,
// but sensitive info such as user access token get cleared
func (c *Client) SafeClone() *Client {
	ret := &Client{
		Credential: ClientCredential{
			ClientID:     c.Credential.ClientID,
			ClientSecret: c.Credential.ClientSecret,
		},
		RedirectUri: c.RedirectUri,
	}
	return ret
}
