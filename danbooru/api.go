package danbooru

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ayame-git/gobooru/danbooru/auth"
)

type DanbooruApi struct {
	Booru      string
	authClient auth.AuthClient
}

func NewDanbooruApi(booru string) (api DanbooruApi) {
	api.Booru = booru
	api.authClient.Client = http.DefaultClient
	return api
}

func (api *DanbooruApi) SetCredentials(login, apiKey string) {
	api.authClient.SetCredentials(login, apiKey)
}

//generic get fucntion
func (api *DanbooruApi) get(url string, v url.Values, data interface{}) error {
	resp, err := api.authClient.Get(api.Booru+url, v)
	if err != nil {
		return err
	}

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
