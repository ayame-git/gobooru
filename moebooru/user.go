package moebooru

import (
	"net/url"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Blacklist []string `json:"blacklisted_tags"`
}

func (api MoebooruApi) GetUsers(v url.Values) (users []User, err error) {
	v = cleanValues(v)
	err = api.get("/user.json", v, &users)
	return
}
