package main

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func main() {
	route()
	start(router)
}

func start(router *gin.Engine) {
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
