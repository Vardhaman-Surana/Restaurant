package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vds/Restraunt/pkg/controller"
	"github.com/vds/Restraunt/pkg/database"
)

type Router struct {
	db database.Database
}

func NewRouter(db database.Database) (*Router, error) {
	router := new(Router)
	router.db = db
	return router,nil
}

func (r *Router) Create(port string) error {
	ginRouter := gin.Default()
	//admin:=ginRouter.Group("/admin/:name")
	//superAdmin:=ginRouter.Group("/superAdmin/:name")



	// create routes
	helloWorld := controller.NewHelloController(r.db)
	login:=controller.NewLogInController(r.db)
	register:=controller.NewRegisterController(r.db)
	ginRouter.GET("/", helloWorld.SayHello)
	ginRouter.POST("/admin/login",login.LogIn)
	ginRouter.POST("/superAdmin/login",login.LogIn)
	ginRouter.POST("/register",register.Register)
	return ginRouter.Run(port)
}
