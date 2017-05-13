package moebooru

import (
	"net/url"
)

func (api MoebooruApi) GetPosts(v url.Values) (posts []Post, err error) {
	v = cleanValues(v)
	err = api.get("/post.json", v, &posts)
	return
}

func (api MoebooruApi) CreatePost(v url.Values) (err error) {
	v = cleanValues(v)
	err = api.post("/post/create.json", v, "post")
	return
}
