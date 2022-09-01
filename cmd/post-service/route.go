package main

import (
	"github.com/Vitaly-Baidin/my-rest/internal/post"
	postRepo "github.com/Vitaly-Baidin/my-rest/internal/post/db"
	"github.com/Vitaly-Baidin/my-rest/internal/postlike"
	postlikeRepo "github.com/Vitaly-Baidin/my-rest/internal/postlike/db"
	"github.com/Vitaly-Baidin/my-rest/internal/user"
	userRepo "github.com/Vitaly-Baidin/my-rest/internal/user/db"
	"github.com/Vitaly-Baidin/my-rest/pkg/client/postgresql"
)

func route() {
	client := postgresql.NewPostgresqlClient()

	userRepository := userRepo.NewPostgresqlRepository(client)
	postRepository := postRepo.NewPostgresqlRepository(client, userRepository)
	postlikeRepository := postlikeRepo.NewPostgresqlRepository(client)

	logger.Info("register user router")
	userHandler := user.NewHandler(userRepository)
	userHandler.Register(router)

	logger.Info("register post router")
	postHandler := post.NewHandler(postRepository)
	postHandler.Register(router)

	logger.Info("register postlike router")
	postlikeHandler := postlike.NewHandler(postlikeRepository)
	postlikeHandler.Register(router)
}
