package moebooru

import (
	"net/url"
)

func (api MoebooruApi) GetPosts(v url.Values) (posts []Post, err error) {
	err = api.get("/post.json", v, &posts)
	return
}
