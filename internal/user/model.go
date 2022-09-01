package user

type User struct {
	UserId   int64
	Username string
}

type Request struct {
	Username string
}

type Response struct {
	UserId   int64  `json:"id"`
	Username string `json:"username"`
}
