package moebooru

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ayame-git/gobooru/auth/moebooru"
	"github.com/ayame-git/gobooru/model/moebooru"
)

//TODO: request constatns (e.g. /post.json artist.json)

type API struct {
	Booru       string
	AuthHandler moebooruAuth.AuthHandler
}

//TODO: fix request over https
func (api *API) GetPosts(preq moebooruModel.PostsRequest) (posts []moebooruModel.Posts, err error) {
	data, err := json.Marshal(preq)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", api.Booru+"/post.json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.AuthHandler.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&posts)

	return
}
