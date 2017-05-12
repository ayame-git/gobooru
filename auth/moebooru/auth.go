package moebooruAuth

import (
	"crypto/sha1"
	"net/http"
	"net/http/cookiejar"
)

const salt string = "So-I-Heard-You-Like-Mupkids-?"

type AuthHandler struct {
	Client http.Client
	IsAuth bool
}

func NewAuthHandler(booru, login, password string) (auth AuthHandler, err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}

	auth.Client.Jar = jar

	data := []byte(salt + "--" + password + "--")
	bytes_hash := sha1.Sum(data)
	password_hash := string(bytes_hash[:])

	req, err := http.NewRequest("POST", booru, nil)
	req.SetBasicAuth(login, password_hash)

	resp, err := auth.Client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		auth.IsAuth = true
	}
	return
}
