package handlers

import (
	"net/http"
	"qtmd-php/database"
	"qtmd-php/model"
	"qtmd-php/util"
	"qtmd-php/valid"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u User) Create(c *gin.Context) {
	if err := c.ShouldBind(&valid.UserReq{}); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}
	user := &model.User{
		Name: c.PostForm("name"),
	}

	database.DbEngine().Create(&user)

	util.REST(c, gin.H{
		"data": user,
	})
}

func (u User) Get(c *gin.Context) {
	if err := c.ShouldBind(&valid.UserReq{}); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
	}

	user := &model.User{}

	database.DbEngine().Where("id=?", c.Query("id")).Find(&user)

	util.REST(c, gin.H{
		"data": user,
	})
}

func (u User) Delete(c *gin.Context) {
	if err := c.ShouldBind(&valid.UserReq{}); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}
	user := &model.User{}

	database.DbEngine().Where("id=?", c.PostForm("id")).Delete(&user)

	util.REST(c, gin.H{})
}
