package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vds/Restraunt/pkg/database"
	"github.com/vds/Restraunt/pkg/models"
	"net/http"
)

type RegisterController struct{
	db database.Database
}

func NewRegisterController(db database.Database) *RegisterController{
	registerController:=new(RegisterController)
	registerController.db=db
	return registerController
}

func (r  *RegisterController)Register(c *gin.Context){
	var admin models.Admin
	err:=c.ShouldBindJSON(&admin)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err=r.db.CreateUser(&admin)
}