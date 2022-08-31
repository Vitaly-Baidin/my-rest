package response

type Post struct {
	PostId int64      `json:"postid"`
	UserId int64      `json:"-"`
	Title  string     `json:"title"`
	Likes  []PostLike `json:"likes"`
}
