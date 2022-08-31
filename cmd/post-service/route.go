package main

import (
	"github.com/Vitaly-Baidin/my-rest/internal/service"
)

const (
	prefix      = "/api/v1"
	allUsersURI = prefix + "/users"
	userURI     = prefix + "/users/:id"
)

func route() {

	logger.Info("register user router")
	router.GET(allUsersURI, service.FindAllUsers)
	router.POST(allUsersURI, service.CreateUser)
	router.GET(userURI, service.FindUserById)
	router.PATCH(userURI, service.UpdateUserById)
	router.DELETE(userURI, service.DeleteUserById)

	logger.Info("register post router")
	router.POST("/post/:id", service.CreatePost)
	router.GET("/post/:id", service.GetPostById)
	router.PATCH("/post/:id", service.UpdatePost)
	router.DELETE("/post/:id", service.DeletePost)

	logger.Info("register postlike router")
	router.PATCH("/postlike/:userid/:postid", service.SetPostlike)
}
