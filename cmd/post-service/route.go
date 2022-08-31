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

	//TODO: refactoring POST POSTLIKE USERLIKE

	// USER
	router.GET(allUsersURI, service.FindAllUsers)
	router.POST(allUsersURI, service.CreateUser)
	router.GET(userURI, service.FindUserById)
	router.PATCH(userURI, service.UpdateUserById)
	router.DELETE(userURI, service.DeleteUserById)

	// POST
	// NORMALLY FOR POST `id` goes from JWT but I didn't implement it, it goes from route
	router.POST("/post/:id", service.CreatePost)
	router.GET("/post/:id", service.GetPostById)
	router.PATCH("/post/:id", service.UpdatePost)
	router.DELETE("/post/:id", service.DeletePost)

	// POSTLIKES
	router.PATCH("/postlike/:userid/:postid", service.SetPostlike)
}
