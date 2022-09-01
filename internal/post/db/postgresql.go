package db

import (
	"database/sql"
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/post"
	"github.com/Vitaly-Baidin/my-rest/internal/user"
	"github.com/Vitaly-Baidin/my-rest/pkg/client/postgresql"
)

const (
	getAllPostsByUserIdQuery = "SELECT PostId, Title FROM Posts WHERE UserId = $1;"
	getPostByIdQuery         = "SELECT PostId, Title FROM Posts WHERE PostId = $1;"
	createPostQuery          = "INSERT INTO Posts (UserId, Title) VALUES ($1, $2);"
	updatePostByIdQuery      = "UPDATE Posts SET Title = $1 WHERE PostId = $2;"
	deletePostByIdQuery      = "DELETE FROM Posts WHERE PostId = $1;"
)

type postgresqlRepository struct {
	client         *sql.DB
	userRepository user.Repository
}

func NewPostgresqlRepository(client *sql.DB, userRepository user.Repository) post.Repository {
	return &postgresqlRepository{
		client:         client,
		userRepository: userRepository,
	}
}

func (r *postgresqlRepository) GetAllPostsByUserId(id int64) ([]post.Response, error) {
	var posts []post.Response

	rows, err := r.client.Query(getAllPostsByUserIdQuery, id)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var response post.Response
		if err = rows.Scan(&response.PostId, &response.Title); err != nil {
			return posts, err
		}

		posts = append(posts, response)
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *postgresqlRepository) GetPostById(id int64) (post.Response, error) {
	var postResponse post.Response

	if err := postgresql.Client.QueryRow(getPostByIdQuery, id).Scan(&postResponse.PostId, &postResponse.Title); err != nil {
		if err == sql.ErrNoRows {
			return post.Response{}, fmt.Errorf("no post with id of: %d", id)
		} else {
			return post.Response{}, err
		}
	}

	return postResponse, nil
}

func (r *postgresqlRepository) CreatePost(userId int64, input post.Request) (int64, error) {
	if _, err := r.userRepository.FindUserById(userId); err != nil {
		return 0, err
	}

	var postId int64

	err := r.client.QueryRow(createPostQuery, userId, input.Title).Scan(&postId)
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (r *postgresqlRepository) UpdatePost(id int64, input post.Request) error {
	if _, err := r.GetPostById(id); err != nil {
		return err
	}

	result, err := r.client.Exec(updatePostByIdQuery, input.Title, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}

func (r *postgresqlRepository) DeletePost(id int64) error {
	if _, err := r.GetPostById(id); err != nil {
		return err
	}

	result, err := r.client.Exec(deletePostByIdQuery, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}

	return nil
}
