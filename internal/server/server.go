package server

import (
	"github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(handler.AuthMiddleware())
	e.GET("/login", handler.Login)
	e.POST("/insert", handler.Insert)
	e.POST("/updateBN/:name", handler.UpdateByName)
	e.POST("/updateBI/:id", handler.UpdateById)
	e.GET("/getBN/:name", handler.GetByName)
	e.GET("/getBI/:id", handler.GetById)
	e.GET("/removeBI/:id", handler.RemoveById)
	e.GET("/removeBN/:name", handler.RemoveByName)
	return e
}
