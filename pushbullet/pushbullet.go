package pushbullet

const (
	API_BASE  string = "https://api.pushbullet.com/"
	SITE_BASE string = "https://www.pushbullet.com/"
)

type ClientCredential struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type Client struct {
	Credential  ClientCredential
	RedirectUri string
	token       string
}

func NewClient(clientID, clientSecret, redirectURI string) *Client {
	return &Client{
		Credential: ClientCredential{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		RedirectUri: redirectURI,
	}
}
