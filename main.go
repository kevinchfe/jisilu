package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qtmd-php/routers"
	//"qtmd-php/task"
)

// @title Swagger 演示 API
// @version 1.0
// @description 这是个演示demo.

// @host localhost:8088
// @BasePath
// @Schemes http https

var r *gin.Engine = routers.NewRouter()

func main() {
	//ti := int64(1605754087)
	////ti,_ := strconv.ParseInt("1605754087",10, 64)
	//fmt.Println(time.Unix(ti,0).Format("2006-01-02 15:04:05"))
	//return

	// 定时任务
	//go task.AddCron()

	// 路由
	if err := r.Run(":8088"); err != nil {
		fmt.Println(err.Error())
	}
}
