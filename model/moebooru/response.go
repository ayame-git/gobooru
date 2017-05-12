package moebooruModel

type Posts struct {
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

type Tags struct {
	ID        int    `json:"id"`
	Name      string `json:"anem"`
	Count     int    `json:"count"`
	Type      int    `json:"type"`
	Ambiguous bool   `json:"ambiguous"`
}

type Artists struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	AliasID int      `json:"alias_id"`
	GroupID int      `json:"group_id"`
	URLs    []string `json:"urls"`
}
