package main

import (
	"net/http"

	"AnGO/models"
	"AnGO/util"

	"github.com/gin-gonic/gin"
)

func main() {
	db := util.GetConnection()
	app := gin.New()
	models.IntialData(db)
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	app.Run()
}
