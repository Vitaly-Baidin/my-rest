package db

import (
	"database/sql"
	"errors"
	postRepo "github.com/Vitaly-Baidin/my-rest/internal/post/db"
	"github.com/Vitaly-Baidin/my-rest/internal/postlike"
	userRepo "github.com/Vitaly-Baidin/my-rest/internal/user/db"
)

const (
	getPostlikesByPostIdQuery         = "SELECT UserName FROM PostLikes JOIN Users ON PostLikes.UserId = Users.UserId WHERE PostLikes.PostId = $1;"
	getPostlikeByUserIdAndPostIdQuery = "SELECT COUNT(*) FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
	existsPostlikeQuery               = "SELECT COUNT(*) FROM Posts WHERE UserId = $1 AND PostId = $2;"
	addPostlikeQuery                  = "INSERT INTO PostLikes (UserId, PostId) VALUES ($1, $2);"
	removePostlikeQuery               = "DELETE FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
)

type postgresqlRepository struct {
	client *sql.DB
}

func NewPostgresqlRepository(client *sql.DB) postlike.Repository {
	return &postgresqlRepository{
		client: client,
	}
}

func (r *postgresqlRepository) SetPostlike(userId int64, postId int64) error {
	userRepository := userRepo.NewPostgresqlRepository(r.client)
	postRepository := postRepo.NewPostgresqlRepository(r.client, userRepository)

	if _, err := userRepository.FindUserById(userId); err != nil {
		return err
	}

	if _, err := postRepository.GetPostById(postId); err != nil {
		return err
	}

	exists, err := r.checkIfUserLikesHisOwnPost(userId, postId)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("you cannot like your own posts")
	}

	exists, err = r.checkPostlikeExists(userId, postId)
	if err != nil {
		return err
	}

	var command string
	if exists {
		command = removePostlikeQuery
	} else {
		command = addPostlikeQuery
	}

	result, err := r.client.Exec(command, userId, postId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}

func (r *postgresqlRepository) GetPostlikesByPostId(postId int64) ([]postlike.Postlike, error) {
	var postlikes []postlike.Postlike

	rows, err := r.client.Query(getPostlikesByPostIdQuery, postId)
	if err != nil {
		return postlikes, err
	}
	defer rows.Close()

	for rows.Next() {
		var p postlike.Postlike
		if err := rows.Scan(&p.Username); err != nil {
			return postlikes, err
		}
		postlikes = append(postlikes, p)
	}

	if err = rows.Err(); err != nil {
		return postlikes, err
	}

	return postlikes, nil
}

func (r *postgresqlRepository) checkPostlikeExists(userId int64, postId int64) (bool, error) {
	var ifExists bool

	if err := r.client.QueryRow(getPostlikeByUserIdAndPostIdQuery, userId, postId).Scan(&ifExists); err != nil {
		return ifExists, err
	}

	return ifExists, nil
}

func (r *postgresqlRepository) checkIfUserLikesHisOwnPost(userId int64, postId int64) (bool, error) {
	var ifExists bool

	if err := r.client.QueryRow(existsPostlikeQuery, userId, postId).Scan(&ifExists); err != nil {
		return ifExists, err
	}

	return ifExists, nil
}
