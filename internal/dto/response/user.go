package response

type User struct {
	UserId   int64  `json:"id"`
	Username string `json:"username"`
	Posts    []Post `json:"posts"`
}
