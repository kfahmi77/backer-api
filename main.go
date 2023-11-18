package main

import (
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
	userInput := user.RegisterUserInput{}
	userInput.Name = "Test simpan dari service"
	userInput.Email = "Test email dari service"
	userInput.Occupation = "Test occupation dari service"
	userInput.Password = "Test passwores dari service"

	userService.RegisterUser(userInput)
}
