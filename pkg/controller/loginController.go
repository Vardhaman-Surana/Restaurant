package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vds/Restraunt/pkg/database"
	"github.com/vds/Restraunt/pkg/models"
	"net/http"
	"strings"
)

type LogInController struct{
	db database.Database
}

func NewLogInController(db database.Database) *LogInController{
	logInController:=new(LogInController)
	logInController.db=db
	return logInController
}

func (l *LogInController) LogIn(c *gin.Context){
	isSuper:=0
	if strings.Contains(c.Request.URL.String(),"superAdmin"){
		isSuper=1
	}

	var cred models.Credentials
	err:=c.ShouldBindJSON(&cred)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err =l.db.LogIn(&cred,isSuper)
	if err!=nil{
		fmt.Print(err)
		return
	}
	fmt.Printf("loginSuccessful %v",isSuper)
}

