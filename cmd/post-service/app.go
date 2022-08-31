package main

import (
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/config"
	"github.com/Vitaly-Baidin/my-rest/pkg/logging"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	logger = logging.GetLogger()
)

func main() {
	cfg := config.GetConfig()
	route()
	start(router, cfg)
}

func start(router *gin.Engine, cfg *config.Config) {
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	err := router.Run(fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		panic(err)
	}
}
