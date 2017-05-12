package main

import (
	"github.com/ayame-git/gobooru/api/moebooru"
)

func main() {
	api := moebooru.API{Booru: "http://konachan.com"}
	req := moebooru.NewPostCreateRequest()

	err := api.CreatePost(req)
	if err != nil {
		panic(err)
	}
}
