package storage

import "github.com/labstack/echo"

type (
	// Searcher defines persistent data getters
	Searcher interface {
		Get(c echo.Context, userID string) (*User, error)
		// Get returns user instance by primary key, or error happen during query

		GetByPushbulletID(c echo.Context, pushbulletID string) (*User, error)
		// GetByPushbulletID returns user with specified pushbullet user ident, or error if nothing found

		GetBySecret(c echo.Context, secret string, isAdminSecret bool) (*User, error)
		// GetBySecret returns user with specified simple push secret or admin
		// secret according to isAdminSecret set or not

	}

	// Keeper defines persistent data savers
	Keeper interface {
		NewUser(c echo.Context, user *User) (*User, error)
		// NewUser saves a user into persistent storage.
		// The pushbullet ID of specified user must be checked against duplication.
		// The userID, SimplePushSecret, AdminSecret in input user instance
		// must be ignored and regenerated safely before save.
		// The tokens in incoming instance will be cleared and never be saved.
		// One should call IssueToken to add new tokens to a user after
		// successful creation.
		// Other fields of incoming user instance will be saved as is, or
		// filled automatically if zero value offered.
		// The returned user instance contains final saved values.
		// There is no guarantee that returned user is newly created, that is,
		// incoming user instance may be modified in place.

		IssueToken(c echo.Context, user *User, AccessToken string) (*User, error)
		// IssueToken add a token to specified user, identified by user's
		// primary key (UserID).
		// Token's issue time field will be set to current system time if zero
		// value offered.
		// Other information of the user will not be updated, even though
		// the stored user data might be different with incoming instance.
		// Error will be returned if no such user, or other error encountered.
	}

	// Backend embedded Searcher and Keeper
	Backend interface {
		Searcher
		Keeper
	}
)
