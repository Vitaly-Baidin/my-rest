package db

import (
	"database/sql"
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/user"
)

const (
	findAllUsersQuery   = "SELECT UserId, Username FROM Users;"
	findUserByIdQuery   = "SELECT UserId, Username FROM Users WHERE UserId = $1;"
	createUserQuery     = "INSERT INTO Users (Username) VALUES ($1) RETURNING UserId;"
	updateUserByIdQuery = "UPDATE Users SET Username = $1 WHERE UserId = $2;"
	deleteUserByIdQuery = "DELETE FROM Users WHERE UserId = $1;"
)

type postgresqlRepository struct {
	client *sql.DB
}

func NewPostgresqlRepository(client *sql.DB) user.Repository {
	return &postgresqlRepository{
		client: client,
	}
}

func (r *postgresqlRepository) FindAllUsers() ([]user.Response, error) {
	var users []user.Response

	rows, err := r.client.Query(findAllUsersQuery)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var response user.Response

		if err = rows.Scan(&response.UserId, &response.Username); err != nil {
			return users, err
		}

		users = append(users, response)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}
func (r *postgresqlRepository) FindUserById(id int64) (user.Response, error) {
	var userResponse user.Response

	if err := r.client.QueryRow(findUserByIdQuery, id).Scan(&userResponse.UserId, &userResponse.Username); err != nil {
		if err == sql.ErrNoRows {
			return user.Response{}, fmt.Errorf("no user with id: %d", id)
		} else {
			return user.Response{}, err
		}
	}

	return userResponse, nil

}
func (r *postgresqlRepository) CreateUser(input user.Request) (int64, error) {
	var userId int64
	err := r.client.QueryRow(createUserQuery, input.Username).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil

}
func (r *postgresqlRepository) UpdateUserById(id int64, input user.Request) error {
	if _, err := r.FindUserById(id); err != nil {
		return err
	}

	result, err := r.client.Exec(updateUserByIdQuery, input.Username, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}
	return nil

}
func (r *postgresqlRepository) DeleteUserById(id int64) error {
	if _, err := r.FindUserById(id); err != nil {
		return err
	}

	result, err := r.client.Exec(deleteUserByIdQuery, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}
	return nil

}
