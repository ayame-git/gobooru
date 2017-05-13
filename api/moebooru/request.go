package moebooru

type PostListRequest struct {
	Limit int      `json:"limit,omitempty"`
	Page  int      `json:"page,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type TagListRequest struct {
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
	Order       string `json:"order"`
	ID          int    `json:"id"`
	AfterID     int    `json:"after_id"`
	Name        string `json:"name"`
	NamePattern string `json:"name_pattern"`
}

type ArtistListRequest struct {
	Name  string `json:"name"`
	Order string `json:"order"`
	Page  int    `json:"page"`
}
