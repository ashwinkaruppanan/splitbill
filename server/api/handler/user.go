package handler

import (
	"net/http"
	"splitbill/internal/model"
	"splitbill/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service model.UserServiceInterface
}

func NewUserHandler(service model.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.CreateUser(c, &user)
	if err != nil {
		utils.JSONError(c, err)
		return
	}

	//set cookie here with res.user_id

	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.LoginUser(c, &user)
	if err != nil {
		utils.JSONError(c, err)
		return
	}

	//set cookie here with user_id

	c.JSON(http.StatusOK, res)
}
