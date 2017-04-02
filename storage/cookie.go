package storage

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
)

const (
	cookieUserPrefix string = "U_"
)

var (
	keyDerivationSalt = [][]byte{
		[]byte("SecretlyLeaveItAl0ne"),
		[]byte("AVerySeriousSecret!"),
	}
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

func (b CookieBackend) iterDecodedCookie(c echo.Context, prefix string, do func(cookie *http.Cookie, user *User) (goon bool)) {
	var dst User
	for _, cookie := range c.Cookies() {
		if prefix != "" && !strings.HasPrefix(cookie.Name, prefix) {
			continue
		}
		if err := b.Decode(cookie.Name, cookie.Value, &dst); err != nil {
			continue
		} else {
			fmt.Printf("iter %v, value %s\n", cookie.Name, dst)
			if !do(cookie, &dst) {
				break
			}
		}
	}
}

func (b CookieBackend) foundFirstMatch(c echo.Context, matcher func(u *User) bool, err error) (*User, error) {
	var u *User
	do := func(cookie *http.Cookie, user *User) bool {
		if matcher(user) {
			u = user
			return false
		}
		return true
	}
	b.iterDecodedCookie(c, cookieUserPrefix, do)
	if u == nil {
		return nil, err
	} else {
		return u, nil
	}
}

func (b CookieBackend) userCookieName(userID string) string {
	return cookieUserPrefix + userID
}

func (b CookieBackend) Get(c echo.Context, userID string) (*User, error) {
	u := &User{}
	if err := b.readCookie(c, b.userCookieName(userID), u); err != nil {
		return nil, errIDNotFound
	} else {
		return u, nil
	}
}

func (b CookieBackend) GetByPushbulletID(c echo.Context, pushbulletID string) (*User, error) {
	matcher := func(u *User) bool {
		return u.PushbulletID == pushbulletID
	}
	return b.foundFirstMatch(c, matcher, errPushbulletIDNotFound)
}

func (b CookieBackend) GetBySecret(c echo.Context, secret string, isAdminSecret bool) (*User, error) {
	matcher := func(u *User) bool {
		return (isAdminSecret && secret == u.AdminSecret) || (!isAdminSecret && secret == u.SimplePushSecret)
	}
	return b.foundFirstMatch(c, matcher, errSecretNotFound)
}

func (b CookieBackend) NewUser(c echo.Context, user *User) (*User, error) {
	user.Tokens = make([]Token, 0)
	user.UserID = safeSecret(8)
	user.SimplePushSecret = safeSecret(8)
	user.AdminSecret = safeSecret(8)
	user.CreatedAt = timestamp()
	user.LastSeen = timestamp()
	if err := b.writeCookie(c, b.userCookieName(user.UserID), user); err != nil {
		return user, err
	}
	return user, nil
}

func (b CookieBackend) IssueToken(c echo.Context, user *User, AccessToken string) (*User, error) {
	user.Tokens = make([]Token, 1)
	user.Tokens[0].AccessToken = AccessToken
	user.Tokens[0].IssuedAt = timestamp()
	if err := b.writeCookie(c, b.userCookieName(user.UserID), user); err != nil {
		return user, err
	}
	return user, nil
}
