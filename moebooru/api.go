package moebooru

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ayame-git/gobooru/moebooru/auth"
)

type MoebooruApi struct {
	Booru      string
	authClient auth.AuthClient
}

func NewMoebooruApi(booru string) (api MoebooruApi) {
	api.Booru = booru
	api.authClient.Client = http.DefaultClient
	return api
}

func (api *MoebooruApi) SetCredentials(login, password string) {
	api.authClient.SetCredentials(login, password)
}

//generic get fucntion
func (api *MoebooruApi) get(url string, v url.Values, data interface{}) error {
	resp, err := api.authClient.Get(api.Booru+url, v)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return decodeResp(resp, data)
}

//TODO: parse responses
func decodeResp(resp *http.Response, data interface{}) (err error) {
	return json.NewDecoder(resp.Body).Decode(data)
}

func cleanValues(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}
