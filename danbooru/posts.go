package danbooru

import (
	"net/url"
)

func (api DanbooruApi) GetPosts(v url.Values) (posts []Post, err error) {
	v = cleanValues(v)
	err = api.get("/posts.json", v, &posts)
	return
}

func (api DanbooruApi) GetRandomPost(v url.Values) (post Post, err error) {
	v = cleanValues(v)
	err = api.get("/posts/random.json", v, &post)
	return
}
