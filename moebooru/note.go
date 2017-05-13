package moebooru

import (
	"net/url"
)

type Note struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Body   string `json:"body"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatorID int    `json:"creator_id"`

	IsActive bool `json:"is_active"`
	Version  int  `json:"version"`
}

func (api MoebooruApi) GetNotes(v url.Values) (notes []Note, err error) {
	v = cleanValues(v)
	err = api.get("/note.json", v, &notes)
	return
}
