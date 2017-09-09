package moebooru

import (
	"net/url"
)

func (api MoebooruApi) GetPosts(v url.Values) (posts []Post, err error) {
	v = cleanValues(v)
	err = api.get("/post.json", v, &posts)
	return
}

func (api MoebooruApi) GetRandomPost(v url.Values) (post Post, err error) {
	v = cleanValues(v)
	v.Add("tags", "order:random")
	posts, err := api.GetPosts(v)
	if err != nil {
		return
	}
	post = posts[0]
	return
}
