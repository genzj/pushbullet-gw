package pushbullet

type (
	PushRequest struct {
		Type     string `json:"type"`
		DeviceID string `json:"device_iden"`
		Body     string `json:"body"`
	}

	Note struct {
		PushRequest
		Title string `json:"title"`
	}

	Link struct {
		PushRequest
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	PushResponse struct {
		Active                  bool    `json:"active"`
		Body                    string  `json:"body"`
		Created                 float64 `json:"created"`
		Direction               string  `json:"direction"`
		Dismissed               bool    `json:"dismissed"`
		ID                      string  `json:"iden"`
		Modified                float64 `json:"modified"`
		ReceiverEmail           string  `json:"receiver_email"`
		ReceiverEmailNormalized string  `json:"receiver_email_normalized"`
		ReceiverID              string  `json:"receiver_iden"`
		SenderEmail             string  `json:"sender_email"`
		SenderEmailNormalized   string  `json:"sender_email_normalized"`
		SenderID                string  `json:"sender_iden"`
		SenderName              string  `json:"sender_name"`
		Title                   string  `json:"title"`
		Type                    string  `json:"type"`
	}
)

func (c *Client) PushNote(deviceID, title, body string) (*PushResponse, error) {
	return c.push(&Note{
		PushRequest: PushRequest{
			Type:     "note",
			DeviceID: deviceID,
			Body:     body,
		},
		Title: title,
	})
}

func (c *Client) PushLink(deviceID, title, body, url string) (*PushResponse, error) {
	return c.push(&Link{
		PushRequest: PushRequest{
			Type:     "link",
			DeviceID: deviceID,
			Body:     body,
		},
		Title: title,
		URL:   url,
	})
}

func (c *Client) push(v interface{}) (*PushResponse, error) {
	resp, err := c.retrieve("/v2/pushes", v, &PushResponse{})
	return resp.(*PushResponse), err
}
