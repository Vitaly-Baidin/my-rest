package post

type Post struct {
	PostId int64
	UserId int64
	Title  string
}

type Request struct {
	Title string
}

type Response struct {
	PostId int64  `json:"postid"`
	UserId int64  `json:"-"`
	Title  string `json:"title"`
}
