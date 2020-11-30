package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type File struct {

}

func (f File) UploadFile(c *gin.Context) {
	file,_ := c.FormFile("file")
	fmt.Println(file.Filename)
	return
}
