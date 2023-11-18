package handler

import (
	"github.com/gin-gonic/gin"
	"go-backer-api/user"
	"net/http"
)

type userhandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userhandler {
	return &userhandler{userService}
}
func (receiver *userhandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	// map input dari user ke struct
	// structya di passing sebagai param service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	user, err := receiver.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, user)
}
