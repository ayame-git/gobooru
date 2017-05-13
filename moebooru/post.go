package moebooru

import (
	"strings"
	"time"
)

type Post struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Tags      string `json:"tags"`
	CreatorID int    `json:"creator_id"`
	CreatedAt int64  `json:"created_at"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`

	FileURL string `json:"file_url"`
	Source  string `json:"source"`

	Score int `json:"score"`
}

func (p Post) CreatedAtTime() time.Time {
	return time.Unix(p.CreatedAt, 0)
}

func (p Post) GetTags() []string {
	return strings.Split(p.Tags, " ")
}
