package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mld-nj/voteBe/api"
	"github.com/mld-nj/voteBe/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vote?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	
	//跨域
	r.Use(api.Cors())
	r.GET("/channel", func(c *gin.Context) {
		var channel structs.VoteChannel
		channelId:=c.DefaultQuery("channelId","1")
		db.Where("channelId=?",channelId).Find(&channel)
		dJson, err := json.Marshal(channel)
		if err != nil {
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK, string(dJson))
	})
	r.GET("/allChannel", func(c *gin.Context) {
		var channels []structs.VoteChannel
		db.Find(&channels)
		dJson, err := json.Marshal(channels)
		if err != nil {
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK, string(dJson))
	})
	r.GET("/option",func(c *gin.Context) {
		var options []structs.VoteOption
		channelId:=c.DefaultQuery("channelId","1")
		db.Where("channelId=?",channelId).Find(&options)
		dJson, err := json.Marshal(options)
		if err != nil {
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK, string(dJson))
	})
	r.Run()
}