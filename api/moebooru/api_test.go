package moebooru

import (
	"testing"
)

func TestPosts(t *testing.T) {
	api := API{
		Booru: "http://konachan.com",
	}
	req := PostListRequest{}

	resp, err := api.GetPosts(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestNotes(t *testing.T) {
	api := API{
		Booru: "http://konachan.com",
	}
	req := NoteListRequest{}

	resp, err := api.GetNotes(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
