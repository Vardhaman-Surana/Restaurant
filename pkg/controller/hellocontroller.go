package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vds/Restraunt/pkg/database"
	"net/http"
)

type HelloController struct {
	db database.Database
}

func NewHelloController (db database.Database) *HelloController {
	helloController:=new(HelloController)
	helloController.db=db
	return helloController
}

func (h *HelloController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello":"world",
	})
}
