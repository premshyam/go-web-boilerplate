package router

import (
	"github.com/gin-gonic/gin"
	"boilerplate/controllers"
  )

func InitializeRoutes(){
	r := gin.Default()
	r.GET("/ping", controller.Pong)
	r.Run()
}