package storage

import (
	"database/sql"
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/request"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/response"
	"github.com/Vitaly-Baidin/my-rest/pkg/client/postgresql"
)

const (
	getAllPostsByUserIdQuery = "SELECT PostId, Title FROM Posts WHERE UserId = $1;"
	getPostByIdQuery         = "SELECT PostId, Title FROM Posts WHERE PostId = $1;"
	createPostQuery          = "INSERT INTO Posts (UserId, Title) VALUES ($1, $2);"
	updatePostByIdQuery      = "UPDATE Posts SET Title = $1 WHERE PostId = $2;"
	deletePostByIdQuery      = "DELETE FROM Posts WHERE PostId = $1;"
)

func GetAllPostsByUserId(id int64) ([]response.Post, error) {
	var posts []response.Post

	rows, err := postgresql.Client.Query(getAllPostsByUserIdQuery, id)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post response.Post
		if err = rows.Scan(&post.PostId, &post.Title); err != nil {
			return posts, err
		}

		postlikes, err := GetPostlikesByPostId(post.PostId)
		if err != nil {
			return posts, err
		}

		if len(postlikes) == 0 {
			post.Likes = []response.PostLike{}
		} else {
			post.Likes = postlikes
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func GetPostById(id int64) (response.Post, error) {
	var post response.Post

	if err := postgresql.Client.QueryRow(getPostByIdQuery, id).Scan(&post.PostId, &post.Title); err != nil {
		if err == sql.ErrNoRows {
			return response.Post{}, fmt.Errorf("no post with id of: %d", id)
		} else {
			return response.Post{}, err
		}
	}

	postlikes, err := GetPostlikesByPostId(post.PostId)
	if err != nil {
		return post, err
	}

	if len(postlikes) == 0 {
		post.Likes = []response.PostLike{}
	} else {
		post.Likes = postlikes
	}

	return post, nil
}

func CreatePost(userId int64, input request.Post) (int64, error) {
	if _, err := FindUserById(userId); err != nil {
		return 0, err
	}

	var postId int64

	err := postgresql.Client.QueryRow(createPostQuery, userId, input.Title).Scan(&postId)
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func UpdatePost(id int64, input request.Post) error {
	if _, err := GetPostById(id); err != nil {
		return err
	}

	result, err := postgresql.Client.Exec(updatePostByIdQuery, input.Title, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}

func DeletePost(id int64) error {
	if _, err := GetPostById(id); err != nil {
		return err
	}

	result, err := postgresql.Client.Exec(deletePostByIdQuery, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}
