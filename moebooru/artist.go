package moebooru

import (
	"net/url"
)

type Artist struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	AliasID int      `json:"alias_id"`
	GroupID int      `json:"group_id"`
	URLs    []string `json:"urls"`
}

func (api MoebooruApi) GetArtists(v url.Values) (artists []Artist, err error) {
	v = cleanValues(v)
	err = api.get("/artist.json", v, &artists)
	return
}
