package postgresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB

	username = "root"
	password = "rootroot"
	host     = "localhost"
	port     = "54320"
	database = "postgres"
)

func init() {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username,
		password,
		host,
		port,
		database,
	)
	var err error
	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
}
