package moebooru

import (
	"net/url"
)

type Forum struct {
	ID       int `json:"id"`
	ParentID int `json:"parent_id"`

	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`

	UpdatedAt string `json:"updated_at"`

	Title string `json:"title"`
	Body  string `json:"body"`
}

func (api MoebooruApi) GetForum(v url.Values) (forum []Forum, err error) {
	v = cleanValues(v)
	err = api.get("/forum.json", v, &forum)
	return
}
