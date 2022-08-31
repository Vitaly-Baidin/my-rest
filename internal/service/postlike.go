package service

import (
	"github.com/Vitaly-Baidin/my-rest/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetPostlike(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert userid to int64")
		return
	}

	postId, err := strconv.ParseInt(context.Param("postid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert postid to int64")
		return
	}

	if err := storage.SetPostlike(userId, postId); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "success")
}
