package storage

import (
	"fmt"
)

type MemoryBackend map[string]*User

func (m *MemoryBackend) Get(userID string) (*User, error) {
	if u, ok := (*m)[userID]; ok {
		return u, nil
	} else {
		return nil, fmt.Errorf("id not found")
	}
}

func (m *MemoryBackend) GetByPushbulletID(pushbulletID string) (*User, error) {
	if u, ok := (*m)[pushbulletID]; ok {
		return u, nil
	} else {
		return nil, fmt.Errorf("id not found")
	}
}

func (m *MemoryBackend) GetBySecret(secret string, isAdminSecret bool) (*User, error) {
	for _, u := range *m {
		if (isAdminSecret && u.AdminSecret == secret) || (!isAdminSecret && u.SimplePushSecret == secret) {
			return u, nil
		}
	}
	return nil, fmt.Errorf("secret not found")
}

func (m *MemoryBackend) NewUser(user *User) (*User, error) {
	if user.PushbulletID == "" {
		panic("user pushbullet indent not fetched")
	}
	user.Tokens = make([]Token, 0)
	user.UserID = user.PushbulletID
	user.SimplePushSecret = user.UserID
	user.AdminSecret = user.UserID
	(*m)[user.UserID] = user
	return user, nil
}

func (m *MemoryBackend) IssueToken(user *User, AccessToken string) (*User, error) {
	u, err := m.Get(user.UserID)
	if err != nil {
		return u, err
	}
	if u == nil {
		panic("backend return a nil user without error")
	}
	u.Tokens = make([]Token, 1)
	u.Tokens[0] = Token{
		AccessToken: AccessToken,
	}
	return u, nil
}
