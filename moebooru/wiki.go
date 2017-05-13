package moebooru

import (
	"net/url"
)

type Wiki struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UpdaterID int    `json:"updater_id"`

	Locked  bool `json:"locked"`
	Version int  `json:"version"`
}

func (api MoebooruApi) GetWiki(v url.Values) (wiki []Wiki, err error) {
	v = cleanValues(v)
	err = api.get("/wiki.json", v, &wiki)
	return
}
