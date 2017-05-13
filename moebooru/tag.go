package moebooru

import (
	"net/url"
)

type Tag struct {
	ID        int    `json:"id"`
	Name      string `json:"anem"`
	Count     int    `json:"count"`
	Type      int    `json:"type"`
	Ambiguous bool   `json:"ambiguous"`
}

func (api MoebooruApi) GetTags(v url.Values) (tags []Tag, err error) {
	v = cleanValues(v)
	err = api.get("/tag.json", v, &tags)
	return
}
