package danbooru

import (
	"net/url"
)

func (api DanbooruApi) GetPosts(v url.Values) (posts []Post, err error) {
	v = cleanValues(v)
	err = api.get("/posts.json", v, &posts)
	return
}
