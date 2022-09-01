package postlike

import (
	"github.com/Vitaly-Baidin/my-rest/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	prefix      = "/api/v1"
	postlikeURI = prefix + "/postlike/:userid/:postid"
)

type handler struct {
	repository Repository
}

func NewHandler(repository Repository) handlers.Handler {
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.PATCH(postlikeURI, h.SetPostlike)
}

func (h *handler) SetPostlike(context *gin.Context) {
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

	if err := h.repository.SetPostlike(userId, postId); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "success")
}
