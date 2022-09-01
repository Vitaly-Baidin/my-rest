package postlike

type Repository interface {
	SetPostlike(userId int64, postId int64) error
	GetPostlikesByPostId(postId int64) ([]Postlike, error)
}
