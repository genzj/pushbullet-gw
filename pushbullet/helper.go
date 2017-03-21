package pushbullet

import (
	"net/url"
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
