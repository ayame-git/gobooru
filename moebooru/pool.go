package moebooru

import (
	"net/url"
)

type Pool struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID    int    `json:"user_id"`

	IsPublic    bool   `json:"is_public"`
	PostCount   int    `json:"post_count"`
	Description string `json:"description"`
}

func (api MoebooruApi) GetPools(v url.Values) (pools []Pool, err error) {
	v = cleanValues(v)
	err = api.get("/pool.json", v, &pools)
	return
}
