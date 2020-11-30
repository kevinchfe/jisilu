package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"qtmd-php/model"
	"qtmd-php/util"
	"qtmd-php/valid"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

var trans = valid.InitTrans()

func (a Article) Get(c *gin.Context) {
	var article valid.ArticleGet
	if err := c.ShouldBind(&article); err != nil {
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

	c.JSON(http.StatusOK, "handler")
	return
}

func (a Article) Create(c *gin.Context) {
	var article valid.ArticleCreate

	if err := c.ShouldBind(&article); err != nil {
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

	articles := &model.Article{
		Title:   c.PostForm("title"),
		Desc:    c.PostForm("desc"),
		Content: c.PostForm("content"),
	}
	ars, _ := articles.Create()
	util.REST(c, gin.H{"data": ars})
}

func (a Article) List(c *gin.Context) {
	return
}
