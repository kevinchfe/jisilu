package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// REST 返回信息自动根据code插入message
func REST(c *gin.Context, obj gin.H) {

	if obj["data"] == nil {
		obj["data"] = nil
	}
	if obj["code"] == nil {
		obj["code"] = SUCCESS
	}
	if obj["msg"] == nil {
		obj["msg"] = GetMessage(SUCCESS)
	}

	obj["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, obj)
}
