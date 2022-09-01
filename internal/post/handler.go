package post

import (
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	prefix          = "/api/v1"
	postByUserIdURI = prefix + "/users/:id/posts"
	postURI         = prefix + "/posts/:id"
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
	router.POST(postByUserIdURI, h.CreatePost)
	router.GET(postURI, h.GetPostById)
	router.PATCH(postURI, h.UpdatePost)
	router.DELETE(postURI, h.DeletePost)
}

func (h *handler) CreatePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input Request
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	userId, err := h.repository.CreatePost(id, input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.Writer.Header().Add("Location", fmt.Sprintf("/users/%d", userId))
	context.JSON(http.StatusCreated, fmt.Sprintf("post added to user of id: %d", id))
}

func (h *handler) GetPostById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	post, err := h.repository.GetPostById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, post)
}

func (h *handler) UpdatePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	var input Request
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	if err := h.repository.UpdatePost(id, input); err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, "post updated")
}

func (h *handler) DeletePost(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if err := h.repository.DeletePost(id); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusNoContent, "post removed")
}
