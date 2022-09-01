package user

type Repository interface {
	FindAllUsers() ([]Response, error)
	FindUserById(id int64) (Response, error)
	CreateUser(input Request) (int64, error)
	UpdateUserById(id int64, input Request) error
	DeleteUserById(id int64) error
}
