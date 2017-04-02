package storage

import "fmt"

var (
	errIDNotFound           = fmt.Errorf("id not found")
	errPushbulletIDNotFound = fmt.Errorf("Pushbullet id not found")
	errSecretNotFound       = fmt.Errorf("Secret not found")
)
