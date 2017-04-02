package storage

import (
	"github.com/labstack/echo"
)

type MemoryBackend map[string]*User

func (m *MemoryBackend) Get(c echo.Context, userID string) (*User, error) {
	if u, ok := (*m)[userID]; ok {
		return u, nil
	} else {
		return nil, errIDNotFound
	}
}

func (m *MemoryBackend) GetByPushbulletID(c echo.Context, pushbulletID string) (*User, error) {
	if u, ok := (*m)[pushbulletID]; ok {
		return u, nil
	} else {
		return nil, errPushbulletIDNotFound
	}
}

func (m *MemoryBackend) GetBySecret(c echo.Context, secret string, isAdminSecret bool) (*User, error) {
	for _, u := range *m {
		if (isAdminSecret && secret == u.AdminSecret) || (!isAdminSecret && secret == u.SimplePushSecret) {
			return u, nil
		}
	}
	return nil, errSecretNotFound
}

func (m *MemoryBackend) NewUser(c echo.Context, user *User) (*User, error) {
	detectSecret := func(key string) bool {
		u, err := m.GetBySecret(c, key, false)
		if u != nil && err == nil {
			return true
		}
		u, err = m.GetBySecret(c, key, true)
		if u != nil && err == nil {
			return true
		}
		return false
	}
	if user.PushbulletID == "" {
		panic("user pushbullet indent not fetched")
	}
	user.Tokens = make([]Token, 0)
	user.UserID = user.PushbulletID
	user.SimplePushSecret = uniqueSecret(8, detectSecret)
	user.AdminSecret = uniqueSecret(8, detectSecret)
	user.CreatedAt = timestamp()
	user.LastSeen = timestamp()
	(*m)[user.UserID] = user
	return user, nil
}

func (m *MemoryBackend) IssueToken(c echo.Context, user *User, AccessToken string) (*User, error) {
	u, err := m.Get(c, user.UserID)
	if err != nil {
		return u, err
	}
	if u == nil {
		panic("backend return a nil user without error")
	}
	u.Tokens = make([]Token, 1)
	u.Tokens[0] = Token{
		AccessToken: AccessToken,
		IssuedAt:    timestamp(),
	}
	return u, nil
}
