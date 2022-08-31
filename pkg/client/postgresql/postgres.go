package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/config"
	"github.com/Vitaly-Baidin/my-rest/pkg/logging"

	_ "github.com/lib/pq"
)

var Client *sql.DB

func init() {
	cfg := config.GetConfig()
	logger := logging.GetLogger()

	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Storage.Username,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.Database,
	)

	Client, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		logger.Fatal(err)
		panic(err)
	}
}
