package service

import (
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/dto/request"
	"github.com/Vitaly-Baidin/my-rest/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(context *gin.Context) {
	users, err := storage.FindAllUsers()
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusOK, users)
}

func CreateUser(context *gin.Context) {
	var input request.User

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	userId, err := storage.CreateUser(input)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.Writer.Header().Add("Location", fmt.Sprintf("http://localhost:3000/users/%d", userId))
	context.JSON(http.StatusCreated, "user added")
}

func FindUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	user, err := storage.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input request.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if err := storage.UpdateUserById(id, input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "user updated")
}

func DeleteUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if err := storage.DeleteUserById(id); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusNoContent, "user removed")
}
