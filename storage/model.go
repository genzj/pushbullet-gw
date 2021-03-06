package storage

type (
	Token struct {
		AccessToken string `json:"accessToken"`
		IssuedAt    int64  `json:"issuedAt"`
	}

	User struct {
		UserID           string  `json:"userId"`
		Name             string  `json:"name"`
		Email            string  `json:"email"`
		PushbulletID     string  `json:"indent"`
		LastSeen         int64   `json:"lastSeen"`
		CreatedAt        int64   `json:"created"`
		SimplePushSecret string  `json:"simplePushSecret"`
		AdminSecret      string  `json:"adminSecret"`
		Tokens           []Token `json:"tokens"`
	}
)

func (u User) Clone() *User {
	return &User{
		UserID:           u.UserID,
		Name:             u.Name,
		Email:            u.Email,
		PushbulletID:     u.PushbulletID,
		LastSeen:         u.LastSeen,
		CreatedAt:        u.CreatedAt,
		SimplePushSecret: u.SimplePushSecret,
		AdminSecret:      u.AdminSecret,
		Tokens:           u.Tokens,
	}
}
