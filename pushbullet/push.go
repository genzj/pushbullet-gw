package pushbullet

type (
	PushRequest struct {
		Type     string `json:"type"`
		DeviceID string `json:"device_iden"`
	}

	Note struct {
		PushRequest
		Title string `json:"title"`
		Body  string `json:"body"`
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
	resp, err := c.retrieve("/v2/pushes", &Note{
		PushRequest: PushRequest{
			Type:     "note",
			DeviceID: deviceID,
		},
		Title: title,
		Body:  body,
	}, &PushResponse{})
	return resp.(*PushResponse), err
}
