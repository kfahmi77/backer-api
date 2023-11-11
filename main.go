package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-backer-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	//dsn := "root:@tcp(127.0.0.1:3306)/backer?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println("connection success")
	//
	//var users []user.User
	//
	//db.Find(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.Name)
	//	fmt.Println(u.AvatarFileName)
	//
	//}
	router := gin.Default()
	router.GET("/handler", handler)
	err := router.Run(":7071")
	if err != nil {
		return
	}

}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/backer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connection success")

	var users []user.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
