package auth

import (
	"crypto/sha1"
	"encoding/hex"
)

const salt string = "So-I-Heard-You-Like-Mupkids-?"

type Credentials struct {
	Login        string
	PasswordHash string
}

func NewCredentials(login, password string) (c Credentials) {
	b := []byte(salt + password)

	c.Login = login
	c.PasswordHash = hash(b)
	return
}

func hash(b []byte) (sha string) {
	hasher := sha1.New()
	hasher.Write(b)
	sha = hex.EncodeToString(hasher.Sum(nil))
	return
}
