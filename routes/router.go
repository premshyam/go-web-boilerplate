package router

import (
	"github.com/gin-gonic/gin"
	"boilerplate/controllers"
  )

func InitializeRoutes(){
	r := gin.Default()
	controllerRoutes(r)
	r.Run()
}

/**
Seperate route function for seperate controller
All controller specific dependency injection and
controller initialization will happen from there specific function
*/
func controllerRoutes(r *gin.Engine){
	statusController := controller.NewController() 
	r.GET("/ping", statusController.Pong)
}