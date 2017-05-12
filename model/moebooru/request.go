package moebooruModel

type PostListRequest struct {
	Limit int      `json:"limit,omitempty"`
	Page  int      `json:"page,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type PostCreateRequest struct {
	Tags           string `form:"post[tags]"`
	File           string `form:"post[file]"`
	Rating         string `form:"post[rating]"`
	Source         string `form:"post[source]"`
	IsRatingLocked bool   `form:"post[is_rating_locked]"`
	IsNoteLocked   bool   `form:"post[is_note_locked]"`
	ParentID       int    `form:"post[parent_id]"`
	MD5            string `form:"post[md5]"`
}
