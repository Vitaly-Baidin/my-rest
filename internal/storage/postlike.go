package storage

import (
	"errors"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/response"
	"github.com/Vitaly-Baidin/my-rest/pkg/client/postgresql"
)

const (
	getPostlikesByPostIdQuery         = "SELECT UserName FROM PostLikes JOIN Users ON PostLikes.UserId = Users.UserId WHERE PostLikes.PostId = $1;"
	getPostlikeByUserIdAndPostIdQuery = "SELECT COUNT(*) FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
	existsPostlikeQuery               = "SELECT COUNT(*) FROM Posts WHERE UserId = $1 AND PostId = $2;"
	addPostlikeQuery                  = "INSERT INTO PostLikes (UserId, PostId) VALUES ($1, $2);"
	removePostlikeQuery               = "DELETE FROM PostLikes WHERE UserId = $1 AND PostId = $2;"
)

func GetPostlikesByPostId(postId int64) ([]response.PostLike, error) {
	var postlikes []response.PostLike

	rows, err := postgresql.Client.Query(getPostlikesByPostIdQuery, postId)
	if err != nil {
		return postlikes, err
	}
	defer rows.Close()

	for rows.Next() {
		var postlike response.PostLike
		if err := rows.Scan(&postlike.Username); err != nil {
			return postlikes, err
		}
		postlikes = append(postlikes, postlike)
	}

	if err = rows.Err(); err != nil {
		return postlikes, err
	}

	return postlikes, nil
}

func SetPostlike(userId int64, postId int64) error {

	if _, err := FindUserById(userId); err != nil {
		return err
	}

	if _, err := GetPostById(postId); err != nil {
		return err
	}

	exists, err := checkIfUserLikesHisOwnPost(userId, postId)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("you cannot like your own posts")
	}

	exists, err = checkPostlikeExists(userId, postId)
	if err != nil {
		return err
	}

	var command string
	if exists {
		command = removePostlikeQuery
	} else {
		command = addPostlikeQuery
	}

	result, err := postgresql.Client.Exec(command, userId, postId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}

func checkPostlikeExists(userId int64, postId int64) (bool, error) {
	var ifExists bool

	if err := postgresql.Client.QueryRow(getPostlikeByUserIdAndPostIdQuery, userId, postId).Scan(&ifExists); err != nil {
		return ifExists, err
	}

	return ifExists, nil
}

func checkIfUserLikesHisOwnPost(userId int64, postId int64) (bool, error) {
	var ifExists bool

	if err := postgresql.Client.QueryRow(existsPostlikeQuery, userId, postId).Scan(&ifExists); err != nil {
		return ifExists, err
	}

	return ifExists, nil
}
