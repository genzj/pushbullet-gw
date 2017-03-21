package pushbullet

import (
	"fmt"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
)

func (c *Client) ActionURL(action string) *url.URL {
	u, _ := url.Parse(API_BASE)
	u.Path = action
	return u
}

// AuthURL returns link to authorize page the user should be guided to
func (c *Client) AuthURL() *url.URL {
	u, _ := url.Parse(SITE_BASE)
	u.Path = "/authorize"

	v := url.Values{}
	v.Set("client_id", c.Credential.ClientID)
	v.Set("redirect_uri", c.RedirectUri)
	v.Set("response_type", "code")
	u.RawQuery = v.Encode()

	return u
}

func (c *Client) sendRequest(path string, body interface{}) (*resty.Response, error) {
	var resp *resty.Response
	var err error

	if !c.HasToken() {
		log.Errorf("token is required to issue request, LoadToken or RefreshToken before continue")
		return nil, fmt.Errorf("token missing")
	}

	req := resty.R().SetHeader("Access-Token", c.token)

	if body == nil {
		resp, err = req.Get(c.ActionURL(path).String())
	} else {
		resp, err = req.SetBody(body).Post(c.ActionURL(path).String())
	}

	if err != nil {
		log.Error(err)
		return nil, err
	} else if resp.StatusCode()/100 != 2 {
		log.Error(resp.String())
		return nil, fmt.Errorf(resp.String())
	} else {
		log.Info(resp)
		return resp, nil
	}
}
