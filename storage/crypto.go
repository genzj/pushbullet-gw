package storage

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaximumUniqueIteration = 100
)

func safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}

// safeSecret returns a safe random secret base64 string derived from random generator
func safeSecret(len int) string {
	bs := make([]byte, len/4*3+3)
	safeRandom(bs)
	return string(base64.StdEncoding.EncodeToString(bs)[:len])
}

func uniqueSecret(len int, judge func(secret string) bool) string {
	iter := 0
	ret := safeSecret(len)
	for judge(ret) {
		ret = safeSecret(len)
		iter += 1
		if iter > MaximumUniqueIteration {
			panic("cannot find a unique secret, enlarge length and retry")
		}
	}
	return ret
}

func hashKey(key string) string {
	if bs, err := bcrypt.GenerateFromPassword([]byte(key), 10); err != nil {
		panic(err)
	} else {
		return string(bs)
	}
}

func verifyKey(key, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(key))
	return err == nil
}
