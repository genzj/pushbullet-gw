package pushbullet

import (
	"testing"
)

func Test_AuthUrl(t *testing.T) {
	expected := "https://www.pushbullet.com/authorize?client_id=YW7uItOzxPFx8vJ4&redirect_uri=http%3A%2F%2Fwww.catpusher.com%2Fauth_complete&response_type=code"
	c := NewClient(
		"YW7uItOzxPFx8vJ4",
		"MmA98EDg0pjr4fZw",
		"http://www.catpusher.com/auth_complete",
	)
	authURL := c.AuthURL()

	if authURL.String() != expected {
		t.Errorf(
			"returned url: %v not equals to expected: %v\n",
			authURL.String(),
			expected,
		)
	}
}
