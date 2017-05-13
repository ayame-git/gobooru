package moebooru

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

type Comments struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Body   string `json:"body"`

	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`

	CreatedAt string `json:"created_at"`
}

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

type Notes struct {
	ID     int    `json:"id"`
	PostID int    `json:"post_id"`
	Body   string `json:"body"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatorID int    `json:"creator_id"`

	IsActive bool `json:"is_active"`
	Version  int  `json:"version"`
}

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Blacklist []string `json:"blacklisted_tags"`
}

type Forum struct {
	ID       int `json:"id"`
	ParentID int `json:"parent_id"`

	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`

	UpdatedAt string `json:"updated_at"`

	Title string `json:"title"`
	Body  string `json:"body"`
}

type Pools struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID    int    `json:"user_id"`

	IsPublic    bool   `json:"is_public"`
	PostCount   int    `json:"post_count"`
	Description string `json:"description"`
}
