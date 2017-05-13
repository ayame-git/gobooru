package moebooru

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

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

//TODO: move these

func (api MoebooruApi) GetTags(v url.Values) (tags []Tags, err error) {
	err = api.get("/tag.json", v, &tags)
	return
}

func (api MoebooruApi) GetArtists(v url.Values) (artists []Artists, err error) {
	err = api.get("/artist.json", v, &artists)
	return
}

func (api MoebooruApi) GetComment(id int64) (comment Comments, err error) {
	v := url.Values{}
	v.Set("id", strconv.FormatInt(id, 10))

	err = api.get("/comment/show.json", v, &comment)
	return
}

func (api MoebooruApi) GetWiki(v url.Values) (wiki []Wiki, err error) {
	err = api.get("/wiki.json", v, &wiki)
	return
}
func (api MoebooruApi) GetNotes(v url.Values) (notes []Notes, err error) {
	err = api.get("/note.json", v, &notes)
	return
}

func (api MoebooruApi) GetUsers(v url.Values) (users []Users, err error) {
	err = api.get("/user.json", v, &users)
	return
}

func (api MoebooruApi) GetForum(v url.Values) (forum []Forum, err error) {
	err = api.get("/forum.json", v, &forum)
	return
}

func (api MoebooruApi) GetPools(v url.Values) (pools []Pools, err error) {
	err = api.get("/pool.json", v, &pools)
	return
}
