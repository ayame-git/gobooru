package auth

import (
	"net/http"
	"net/url"

	"github.com/ayame-git/gobooru/utils"
)

type AuthClient struct {
	cr     Credentials
	Client *http.Client
}

func (a *AuthClient) SetCredentials(login, apiKey string) {
	a.cr = Credentials{
		Login:  login,
		ApiKey: apiKey,
	}
}

func (a AuthClient) Get(url string, form url.Values) (resp *http.Response, err error) {
	b, err := utils.ToBuffer(form)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url, b)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err = a.Client.Do(req)
	return
}
