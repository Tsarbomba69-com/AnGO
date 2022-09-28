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
	// models.IntialData(db)
	var provinces []models.APIProvince = models.GetProvinces(db)
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, provinces)
	})
	app.Run()
}
