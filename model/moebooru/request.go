package moebooruModel

type PostsRequest struct {
	Limit int      `json:"limit,omitempty"`
	Page  int      `json:"page,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}
