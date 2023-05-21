package router

import (
	"github.com/gin-gonic/gin"
	"boilerplate/controllers"
  )

func InitializeRoutes(){
	r := gin.Default()
	statusControllerRoutes(r)
	r.Run()
}

/**
Seperate route function for seperate controller
All controller specific dependency injection and
controller initialization will happen from there specific function
*/
func statusControllerRoutes(r *gin.Engine){
	statusController := controller.NewStatusController() 
	r.GET("/ping", statusController.Pong)
}