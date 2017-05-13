package moebooru

import (
	"testing"
)

func TestGet(t *testing.T) {
	api := NewMoebooruApi("http://konachan.com")

	_, err := api.GetPosts(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetArtists(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetTags(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetUsers(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetForum(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetWiki(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetNotes(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = api.GetPools(nil)
	if err != nil {
		t.Fatal(err)
	}
}
