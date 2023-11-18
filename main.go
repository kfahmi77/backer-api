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

	userByEmail, err := userRepository.FindByEmail("test@gmail.com")
	if err != nil {
		fmt.Print(err.Error())
	}
	if userByEmail.ID == 0 {
		fmt.Println("user tidak ditemukan")
	} else {
		fmt.Println("user ditemukan", userByEmail.Name)
	}
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("users", userHandler.RegisterUser)

	router.Run(":7071")
}
