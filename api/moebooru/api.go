package moebooru

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//TODO: request constants

const (
	postList    = "/post.json"
	tagList     = "/tag.json"
	artistList  = "/artist.json"
	commentList = "/comment/show.json"
	wikiList    = "/wiki.json"
	noteList    = "/note.json"
	userList    = "/user.json"
	forumList   = "/forum.json"
	poolList    = "/pool.json"
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

//get comment by id ?
func (api *API) GetComments(req CommentListRequest) (comment Comments, err error) {
	err = api.get(commentList, req, &comment)
	return
}

func (api *API) GetWiki(req WikiListRequest) (wiki []Wiki, err error) {
	err = api.get(wikiList, req, &wiki)
	return
}
func (api *API) GetNotes(req NoteListRequest) (notes []Notes, err error) {
	err = api.get(noteList, req, &notes)
	return
}

func (api *API) GetUsers(req UserListRequest) (users []Users, err error) {
	err = api.get(userList, req, &users)
	return
}

func (api *API) GetForum(req ForumListRequest) (forum []Forum, err error) {
	err = api.get(forumList, req, &forum)
	return
}

func (api *API) GetPools(req PoolListRequest) (pools []Pools, err error) {
	err = api.get(poolList, req, &pools)
	return
}
