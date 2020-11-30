package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	_ "qtmd-php/docs"
	"qtmd-php/handlers"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 路由配置表
func NewRouter() *gin.Engine {
	// 开启彩色打印
	gin.ForceConsoleColor()
	//gin.SetMode(gin.ReleaseMode)

	// 记录到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// http 请求路径
		result, _ := http.Get("https://www.shanbay.com/api/v1/vocabtest/vocabularies/?category=CET4")
		// 结果集转为string
		bytes, _ := ioutil.ReadAll(result.Body)

		type choices struct {
			Pk         int64  `json:"pk"`
			Rank       int64  `json:"rand" `
			Definition string `json:"definition"`
		}

		type restData struct {
			Content           string     `json:"content"`
			Pk                int64      `json:"pk"`
			Rank              int64      `json:"rand" `
			DefinitionChoices []*choices `json:"definition_choices"`
		}

		type rest struct {
			Msg  string      `json:"msg"`
			Code int64       `json:"code"`
			Data []*restData `json:"data"`
		}

		var tempMap rest
		// string转为 map
		json.Unmarshal([]byte(bytes), &tempMap)

		fmt.Println(tempMap.Data)

		c.JSON(http.StatusOK, 200)
	})

	// kzz数据整理
	r.GET("/jisilu", handlers.KzzDb)
	r.GET("/yyb", handlers.KzzYyb)

	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, viper.GetString("Database.DBType"))
	})

	article := handlers.Article{}
	user := handlers.User{}
	//file := handlers.File{}

	r.POST("/user/create", user.Create)
	r.GET("/user/get", user.Get)
	r.POST("/user/delete", user.Delete)
	r.GET("/handler", article.Get)
	r.POST("/articles", article.Create)
	r.GET("/articles", article.List)
	//r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(file.Filename)
		}
		fmt.Println(file.Filename)

	})

	//r.GET("/user", user.GetFirst)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

//func upload(w http.ResponseWriter, r *http.Request)  {
//	r.ParseMultipartForm(32 << 20)
//	file,handler,_ := r.FormFile("file")
//	fmt.Println(handler.Filename)
//}

func JsonExit(code int64, message string, data interface{}) map[string]interface{} {
	json := make(map[string]interface{})
	json["code"] = code
	json["message"] = message
	json["data"] = data

	return json
}
