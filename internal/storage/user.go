package storage

import (
	"database/sql"
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/request"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/response"
	"github.com/Vitaly-Baidin/my-rest/pkg/client/postgresql"
)

const (
	findAllUsersQuery   = "SELECT UserId, Username FROM Users;"
	findUserByIdQuery   = "SELECT UserId, Username FROM Users WHERE UserId = $1;"
	createUserQuery     = "INSERT INTO Users (Username) VALUES ($1) RETURNING UserId;"
	updateUserByIdQuery = "UPDATE Users SET Username = $1 WHERE UserId = $2;"
	deleteUserByIdQuery = "DELETE FROM Users WHERE UserId = $1;"
)

func FindAllUsers() ([]response.User, error) {
	var users []response.User

	rows, err := postgresql.Client.Query(findAllUsersQuery)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user response.User

		if err = rows.Scan(&user.UserId, &user.Username); err != nil {
			return users, err
		}

		posts, err := GetAllPostsByUserId(user.UserId)
		if err != nil {
			return users, err
		}

		if len(posts) == 0 {
			user.Posts = []response.Post{}
		} else {
			user.Posts = posts
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func FindUserById(id int64) (response.User, error) {
	var user response.User

	if err := postgresql.Client.QueryRow(findUserByIdQuery, id).Scan(&user.UserId, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return response.User{}, fmt.Errorf("no user with id: %d", id)
		} else {
			return response.User{}, err
		}
	}

	posts, err := GetAllPostsByUserId(user.UserId)
	if err != nil {
		return user, err
	}

	if len(posts) == 0 {
		user.Posts = []response.Post{}
	} else {
		user.Posts = posts
	}

	return user, nil
}

func CreateUser(input request.User) (int64, error) {
	var userId int64
	err := postgresql.Client.QueryRow(createUserQuery, input.Username).Scan(&userId)
	if err != nil {
		return 0, err
	}

	//rows, err := result.RowsAffected()
	//if err != nil || rows == 0 {
	//	return 0, err
	//}
	return userId, nil
}

func UpdateUserById(id int64, input request.User) error {
	if _, err := FindUserById(id); err != nil {
		return err
	}

	result, err := postgresql.Client.Exec(updateUserByIdQuery, input.Username, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}
	return nil
}

func DeleteUserById(id int64) error {
	if _, err := FindUserById(id); err != nil {
		return err
	}

	result, err := postgresql.Client.Exec(deleteUserByIdQuery, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return err
	}
	return nil
}
