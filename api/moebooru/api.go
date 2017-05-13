package moebooru

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//TODO: request constants

const (
	postList   = "/post.json"
	tagList    = "/tag.json"
	artistList = "/artist.json"
)

type API struct {
	Booru       string
	AuthHandler AuthHandler
}

//generic get fucntion
//TODO: fix request over https
func (api *API) get(url string, req interface{}, list interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("GET", api.Booru+url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/json")

	resp, err := api.AuthHandler.Client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(list)

	return err
}

func (api *API) GetPosts(req PostListRequest) (posts []Posts, err error) {
	err = api.get(postList, req, &posts)
	return
}

func (api *API) GetTags(req TagListRequest) (tags []Tags, err error) {
	err = api.get(tagList, req, &tags)
	return
}

func (api *API) GetArtists(req ArtistListRequest) (artists []Artists, err error) {
	err = api.get(artistList, req, &artists)
	return
}

//do we even need these?
func NewPostListRequest() (r PostListRequest) {
	return
}

func NewArtistListRequest() (r ArtistListRequest) {
	return
}

func NewTagListRequest() (r ArtistListRequest) {
	return
}
