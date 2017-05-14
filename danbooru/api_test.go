package danbooru

import (
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	api := NewDanbooruApi("https://danbooru.donmai.us")

	query := url.Values{}
	query.Set("tags", "tokisaki_kurumi")
	query.Set("page", "1")

	posts, err := api.GetPosts(query)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(posts[0].FileURL)
}
