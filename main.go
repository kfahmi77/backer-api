package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-backer-api/handler"
	"go-backer-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/backer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	input := user.LoginUserInput{
		Email:    "admin@gmail.com",
		Password: "123s",
	}
	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
	}
	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("users", userHandler.RegisterUser)

	router.Run(":7071")
}
