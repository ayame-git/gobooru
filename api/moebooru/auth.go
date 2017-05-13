package moebooru

import (
	"net/http"
)

const salt string = "So-I-Heard-You-Like-Mupkids-?"

type AuthHandler struct {
	Client http.Client
	IsAuth bool
}

//TODO: make it actually work?
func NewAuthHandler(booru, login, password string) (auth AuthHandler, err error) {
	return
}
