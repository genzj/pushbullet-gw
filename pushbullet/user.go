package pushbullet

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

func (c *Client) GetUser() (*User, error) {
	resp, err := c.retrieve("/v2/users/me", nil, &User{})
	return resp.(*User), err
}
