package main

import (
	"os"
	"project0/controller"
	"project0/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/page/post", func(c *gin.Context) {
		c.Request.ParseForm() //解析后才能读出来
		data := controller.AppendPost(&c.Request.Form, "./data/")
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
