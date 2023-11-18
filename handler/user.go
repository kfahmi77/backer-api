package handler

import (
	"github.com/gin-gonic/gin"
	"go-backer-api/helper"
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
	//tangkap input dari newUser
	// map input dari newUser ke struct
	// structya di passing sebagai param service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	newUser, err := receiver.userService.RegisterUser(input)
	formatter := user.FormatUser(newUser, "testtoken")
	response := helper.ApiResponse("Account has  been registered", http.StatusOK, "success", formatter)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, response)
}
