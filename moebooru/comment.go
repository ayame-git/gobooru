package moebooru

import (
	"net/url"
	"strconv"
)

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Body   string `json:"body"`

	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`

	CreatedAt string `json:"created_at"`
}

func (api MoebooruApi) GetComment(id int64) (comment Comment, err error) {
	v := url.Values{}
	v.Set("id", strconv.FormatInt(id, 10))

	err = api.get("/comment/show.json", v, &comment)
	return
}
