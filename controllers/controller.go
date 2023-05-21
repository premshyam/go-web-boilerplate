package controller

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
)

type statusController struct {
	baseController
}

func NewController() *statusController {
	return &statusController{}
}

func (ctrl *statusController) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
	"message": "pong",
	})
}