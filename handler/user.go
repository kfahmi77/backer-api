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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Account not registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := receiver.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Account not registered", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "testtoken")
	response := helper.ApiResponse("Account has  been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userhandler) Login(c *gin.Context) {
	//user memasukan input email dan password
	//input ditangkap handler
	//mapping dari inpput user ke input struct
	// passing ke service
	//di service mencari dengan bantuan repository user dengan mencocokan password
	var input user.LoginUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Loggin Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Loggin Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, "token123")
	response := helper.ApiResponse("Suceesfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func CheckEmailAvailable(c *gin.Context) {
	//ada input dari user
	//input email dimapping ke struct input
	//struct input dipassing ke service
	//service akan memanggil repository - email sudah ada atau belum
	// repository - db
}
