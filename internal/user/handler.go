package user

import (
	"fmt"
	"github.com/Vitaly-Baidin/my-rest/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	prefix      = "/api/v1"
	allUsersURI = prefix + "/users"
	userURI     = prefix + "/users/:id"
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
	router.GET(allUsersURI, h.FindAllUsers)
	router.POST(allUsersURI, h.CreateUser)
	router.GET(userURI, h.FindUserById)
	router.PATCH(userURI, h.UpdateUserById)
	router.DELETE(userURI, h.DeleteUserById)
}

func (h *handler) FindAllUsers(context *gin.Context) {
	users, err := h.repository.FindAllUsers()
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
	}

	context.JSON(http.StatusOK, users)
}

func (h *handler) CreateUser(context *gin.Context) {
	var input Request

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, "error binding data")
		return
	}

	userId, err := h.repository.CreateUser(input)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	context.Writer.Header().Add("Location", fmt.Sprintf("/users/%d", userId))
	context.JSON(http.StatusCreated, "user added")
}

func (h *handler) FindUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	user, err := h.repository.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, err.Error())
		return
	}

	context.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUserById(context *gin.Context) {
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

	if err := h.repository.UpdateUserById(id, input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "user updated")
}

func (h *handler) DeleteUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot convert id to int64")
		return
	}

	if err := h.repository.DeleteUserById(id); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusNoContent, "user removed")
}
