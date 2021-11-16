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
	r.POST("/login",func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(200,"hello,%s,密码为:%s",username,password)
	})
	r.GET("/vote",func(c *gin.Context) {
		reqIP := c.ClientIP()
		if reqIP == "::1" {
			reqIP = "127.0.0.1"
		}
		optionId1:=c.DefaultQuery("optionId1","0")
		optionId2:=c.DefaultQuery("optionId2","0")
		optionId3:=c.DefaultQuery("optionId3","0")
		var entity structs.VoteEntity
		entityResult:=db.Where("Ip =?",reqIP).Find(&entity)
		fmt.Println(entityResult.RowsAffected)
		if(entityResult.RowsAffected!=0){
			c.String(500,"你已经投过票了")
			return
		}
		// 通过 `RowsAffected` 得到更新的记录数
		result1:=db.Model(structs.VoteOption{}).Where("optionId = ?", optionId1).Update("count", gorm.Expr("count + ?",1))
		result2:=db.Model(structs.VoteOption{}).Where("optionId = ?", optionId2).Update("count", gorm.Expr("count + ?",1))
		result3:=db.Model(structs.VoteOption{}).Where("optionId = ?", optionId3).Update("count", gorm.Expr("count + ?",1))
		if(result1.Error==nil&&result2.Error==nil&&result3.Error==nil){
			// 条件更新
			user := structs.VoteEntity{Ip: reqIP}

			result := db.Create(&user) // 通过数据的指针来创建
			if(result.Error==nil){
				c.String(200,"投票成功")
			}

		}else{
			c.String(500,"投票失败")
		}
	// result.RowsAffected // 更新的记录数
	// result.Error        // 更新的错误
	})
	r.Run()
}