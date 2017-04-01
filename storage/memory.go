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
		if (isAdminSecret && verifyKey(secret, u.AdminSecret)) || (!isAdminSecret && verifyKey(secret, u.SimplePushSecret)) {
			return u, nil
		}
	}
	return nil, fmt.Errorf("secret not found")
}

func (m *MemoryBackend) NewUser(user *User) (*User, error) {
	detectSecret := func(key string) bool {
		u, err := m.GetBySecret(key, false)
		if u != nil && err == nil {
			return true
		}
		u, err = m.GetBySecret(key, true)
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
	newUser := user.Clone()
	newUser.SimplePushSecret = hashKey(newUser.SimplePushSecret)
	newUser.AdminSecret = hashKey(newUser.AdminSecret)
	(*m)[newUser.UserID] = newUser
	fmt.Println("a", user, newUser)
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
