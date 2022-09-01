package post

type Repository interface {
	GetAllPostsByUserId(id int64) ([]Response, error)
	GetPostById(id int64) (Response, error)
	CreatePost(userId int64, input Request) (int64, error)
	UpdatePost(id int64, input Request) error
	DeletePost(id int64) error
}
