package danbooru

import (
	"strings"
	"time"
)

type Post struct {
	ID        int    `json:"id"`
	Artist    string `json:"tag_string_artist"`
	Tags      string `json:"tags"`
	CreatorID int    `json:"creator_id"`
	CreatedAt string `json:"created_at"`
	Width     int    `json:"image_width"`
	Height    int    `json:"image_height"`

	FileURL string `json:"file_url"`
	Source  string `json:"source"`

	Score int `json:"score"`
}

//doesn't seem to work
func (p Post) CreatedAtTime() (time.Time, error) {
	layout := "2017-05-14T02:10:24.232-04:00"
	return time.Parse(layout, p.CreatedAt)
}

func (p Post) GetTags() []string {
	return strings.Split(p.Tags, " ")
}
