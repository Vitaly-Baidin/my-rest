package service

import (
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/request"
	"github.com/Vitaly-Baidin/my-rest/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input request.Post
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	userId, err := storage.CreatePost(id, input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.Writer.Header().Add("Location", fmt.Sprintf("http://localhost:3000/users/%d", userId))
	context.JSON(http.StatusCreated, fmt.Sprintf("post added to user of id: %d", id))
}

func GetPostById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	post, err := storage.GetPostById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, post)
}

func UpdatePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input request.Post
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if err := storage.UpdatePost(id, input); err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, "post updated")
}

func DeletePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if err := storage.DeletePost(id); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusNoContent, "post removed")
}
