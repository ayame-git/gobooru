package auth

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/ayame-git/gobooru/utils"
)

type AuthClient struct {
	cr     Credentials
	Client *http.Client
}

func (a *AuthClient) SetCredentials(login, password string) {
	a.cr = NewCredentials(login, password)

	cookieJar, _ := cookiejar.New(nil)
	a.Client = &http.Client{
		Jar: cookieJar,
	}

	//login here but idk how
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

func (a AuthClient) Post(url string, v url.Values, field string) (resp *http.Response, err error) {
	//no need if cookies are set
	//v.Set("login", a.cr.Login)
	//v.Set("password_hash", a.cr.PasswordHash)

	b, err := utils.ToForm(v, field)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "multipart/form-data")

	resp, err = a.Client.Do(req)
	return
}
