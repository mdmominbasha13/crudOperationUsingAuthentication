package router

import (
	"github.com/CrudOperationUsingAuthentication/pkg/authentication"
	"github.com/CrudOperationUsingAuthentication/pkg/config"
	"github.com/CrudOperationUsingAuthentication/pkg/controller"
	"github.com/CrudOperationUsingAuthentication/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()

	resource := router.Group("/api")
	router.POST("/login", controller.Login)

	resource.Use(middleware.LogRequestInfo(), authentication.Auth())
	{
		resource.GET("/book/", controller.GetBook)
		resource.GET("/book/:bookid", controller.GetBookById)
		resource.POST("/book/", controller.CreateBook)
		resource.PUT("/book/:bookid", controller.UpdateBookById)
		resource.DELETE("/book/:bookid", controller.DeleteBookById)

	}
	return router
}
