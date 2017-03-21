package pushbullet

import (
	"encoding/json"
)

type (
	User struct {
		ID              string  `json:"iden"`
		Name            string  `json:"name"`
		Email           string  `json:"email"`
		EmailNormalized string  `json:"email_normalized"`
		ImageUrl        string  `json:"image_url"`
		MaxUploadSize   int     `json:"max_upload_size"`
		Created         float64 `json:"created"`
		Modified        float64 `json:"modified"`
	}
)

func (c *Client) Me() (*User, error) {
	u := &User{}
	if resp, err := c.sendRequest("/v2/users/me", nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(resp.Body(), u); err != nil {
		return nil, err
	} else {
		return u, nil
	}
}
