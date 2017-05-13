package moebooru

type PostListRequest struct {
	Limit int      `json:"limit,omitempty"`
	Page  int      `json:"page,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type TagListRequest struct {
	Limit       int    `json:"limit,omitempty"`
	Page        int    `json:"page,omitempty"`
	Order       string `json:"order,omitempty"`
	ID          int    `json:"id,omitempty"`
	AfterID     int    `json:"after_id,omitempty"`
	Name        string `json:"name,omitempty"`
	NamePattern string `json:"name_pattern,omitempty"`
}

type ArtistListRequest struct {
	Name  string `json:"name,omitempty"`
	Order string `json:"order,omitempty"`
	Page  int    `json:"page,omitempty"`
}

type CommentListRequest struct {
	ID int `json:"id,omitempty"`
}

type WikiListRequest struct {
	Order string `json:"order,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Page  int    `json:"page,omitempty"`
	Query string `json:"query,omitempty"`
}

type NoteListRequest struct {
	PostID int `json:"post_id,omitempty"`
}

type UserListRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ForumListRequest struct {
	ParentID int `json:"parent_id,omitempty"`
}

type PoolListRequest struct {
	Query string `json:"query,omitempty"`
	Page  int    `json:"page,omitempty"`
}
