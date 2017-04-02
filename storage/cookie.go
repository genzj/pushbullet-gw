package storage

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/pbkdf2"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

const (
	cookieName string = "U_Pushbullet"
)

var (
	keyDerivationSalt = [][]byte{
		[]byte("SecretlyLeaveItAl0ne"),
		[]byte("AVerySeriousSecret!"),
	}
	errCookieNotSet = fmt.Errorf("Cookie not set")
)

type CookieBackend struct {
	*securecookie.SecureCookie
}

func NewCookieBackend(secret string) *CookieBackend {
	hkb := pbkdf2.Key([]byte(secret), keyDerivationSalt[0], 4096, 32, sha1.New)
	blkb := pbkdf2.Key([]byte(secret), keyDerivationSalt[1], 4096, 32, sha1.New)
	return &CookieBackend{securecookie.New(hkb, blkb)}
}

func (b CookieBackend) readCookie(c echo.Context, name string, dst interface{}) error {
	cookie, err := c.Cookie(name)
	if err != nil {
		return err
	}
	return b.Decode(name, cookie.Value, dst)
}

func (b CookieBackend) writeCookie(c echo.Context, name string, value interface{}) error {
	encoded, err := b.Encode(name, value)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    encoded,
		HttpOnly: true,
		Expires:  time.Now().AddDate(2, 0, 0),
	})
	return nil
}

func (b CookieBackend) Get(c echo.Context, userID string) (*User, error) {
	u := &User{}
	if userID != cookieName {
		return nil, errIDNotFound
	} else if err := b.readCookie(c, cookieName, u); err != nil {
		return nil, errIDNotFound
	} else {
		return u, nil
	}
}

func (b CookieBackend) GetByPushbulletID(c echo.Context, pushbulletID string) (*User, error) {
	u, err := b.Get(c, cookieName)
	if err == nil && u.PushbulletID == pushbulletID {
		return u, nil
	}
	return nil, errPushbulletIDNotFound
}

func (b CookieBackend) GetBySecret(c echo.Context, secret string, isAdminSecret bool) (*User, error) {
	u, err := b.Get(c, cookieName)
	if err == nil && ((isAdminSecret && secret == u.AdminSecret) || (!isAdminSecret && secret == u.SimplePushSecret)) {
		return u, nil
	}
	return nil, errSecretNotFound
}

func (b CookieBackend) NewUser(c echo.Context, user *User) (*User, error) {
	user.Tokens = make([]Token, 0)
	user.UserID = cookieName
	user.SimplePushSecret = cookieName
	user.AdminSecret = cookieName
	user.CreatedAt = timestamp()
	user.LastSeen = timestamp()
	if err := b.writeCookie(c, user.UserID, user); err != nil {
		return user, err
	}
	return user, nil
}

func (b CookieBackend) IssueToken(c echo.Context, user *User, AccessToken string) (*User, error) {
	user.Tokens = make([]Token, 1)
	user.Tokens[0].AccessToken = AccessToken
	user.Tokens[0].IssuedAt = timestamp()
	if err := b.writeCookie(c, user.UserID, user); err != nil {
		return user, err
	}
	return user, nil
}
