package model

import (
	. "qtmd-php/database"
	"qtmd-php/util"

	"github.com/gin-gonic/gin"
)

// User 用户表
// Id 必须为int64 才会返回自增Id
type User struct {
	// 主键
	Id int64 `gorm:"column:id; PRIMARY_KEY" json:"id" form:"id"`
	// 名称
	Name string `gorm:"column:name" json:"name" form:"name" binding:"required`
}

// DB 数据库句柄
// var DB = database.DB

func init() {
	DB.AutoMigrate(&User{})
}

// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 用户信息   //swagger API分类标签, 同一个tag为一组
// @accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
// @Produce  json  //设置返回数据的类型和编码
// @Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param name query string false "name"
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /test/{id} [get]    //路由信息，一定要写上

// @Summary 测试接口
// @Description 描述信息
// @Success 200 {object} util.JSONResult{data=User} "返回结果"
// @Router /user [get]
func (u User) GetFirst(c *gin.Context) {

	DB.First(&u)

	util.REST(c, gin.H{"data": &u})
}
