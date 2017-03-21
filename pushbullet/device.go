package pushbullet

type (
	Device struct {
		ID           string  `json:"iden"`
		Active       bool    `json:"active"`
		AppVersion   int     `json:"app_version"`
		Manufacturer string  `json:"manufacturer"`
		Model        string  `json:"model"`
		Nickname     string  `json:"nickname"`
		PushToken    string  `json:"push_token"`
		Created      float64 `json:"created"`
		Modified     float64 `json:"modified"`
	}

	DeviceList struct {
		Devices []Device `json:"devices"`
	}
)

func (c *Client) ListDevices() (*DeviceList, error) {
	resp, err := c.retrieve("/v2/devices", nil, &DeviceList{})
	return resp.(*DeviceList), err
}
