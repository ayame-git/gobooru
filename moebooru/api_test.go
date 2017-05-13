package moebooru

import (
	"net/url"
	"testing"
)

func TestPosts(t *testing.T) {
	api := NewMoebooruApi("http://konachan.com")
	req := url.Values{}

	resp, err := api.GetPosts(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestNotes(t *testing.T) {
	api := NewMoebooruApi("http://konachan.com")
	req := url.Values{}

	resp, err := api.GetNotes(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
