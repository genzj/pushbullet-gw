package storage

type (
	Token struct {
		AccessToken string `json:"accessToken"`
		IssuedAt    int32  `json:"issuedAt"`
	}

	User struct {
		UserID           string  `json:"userId"`
		Name             string  `json:"name"`
		Email            string  `json:"email"`
		PushbulletID     string  `json:"indent"`
		LastSeen         int32   `json:"lastSeen"`
		CreatedAt        int32   `json:"created"`
		SimplePushSecret string  `json:"simplePushSecret"`
		AdminSecret      string  `json:"adminSecret"`
		Tokens           []Token `json:"tokens"`
	}
)
