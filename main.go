package main

import (
	"go-backer-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/backer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	user := user.User{
		ID:             2,
		Name:           "coba",
		Occupation:     "coba",
		Email:          "coba",
		AvatarFileName: "coba.jpg",
		Role:           "biasa",
		Token:          "teoken123",
		PasswordHash:   "test",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	userRepository.Save(user)
}
