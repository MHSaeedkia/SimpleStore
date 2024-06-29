package main

import (
	"github.com/gin-gonic/gin"
)

func getEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(authMiddleware())
	e.GET("/login", login)
	e.POST("/insert", insert)
	e.POST("/updateBN/:name", updateByName)
	e.POST("/updateBI/:id", updateById)
	e.GET("/getBN/:name", getByName)
	e.GET("/getBI/:id", getById)
	e.GET("/removeBI/:id", removeById)
	e.GET("/removeBN/:name", removeByName)
	return e
}

func main() {
	e := getEngine()
	e.Run()
}
