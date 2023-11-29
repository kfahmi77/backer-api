package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-backer-api/auth"
	"go-backer-api/helper"
	"go-backer-api/user"
	"net/http"
)

type userhandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userhandler {
	return &userhandler{userService, authService}
}
func (h *userhandler) RegisterUser(c *gin.Context) {
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
	// Check if the email is available
	isEmailAvailable, err := h.userService.IsEmailAvailable(user.CheckEmailInput{Email: input.Email})
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Email availability check failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isEmailAvailable {
		// Email is not available; return an error
		errorMessage := gin.H{"errors": "Email has been used"}
		response := helper.ApiResponse("Email has been used", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Account not registered", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Account not registered", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)
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
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ApiResponse("Loggin failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.ApiResponse("Suceesfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userhandler) CheckEmailAvailable(c *gin.Context) {
	//ada input dari user
	//input email dimapping ke struct input
	// dipassing ke service
	// akan memanggil repository - email sudah ada atau belum
	// repository - db
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Email checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email has been used"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
func (h *userhandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload image avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload image avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload image avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Success to upload image avatar", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
	return
}
