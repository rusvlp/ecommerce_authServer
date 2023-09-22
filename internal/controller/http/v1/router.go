package v1

import (
	"ca_kobercams/internal/controller/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserController user.UserController
	Handler        *gin.Engine
}

func (router *Router) ConfigureRouter() {

	v1 := router.Handler.Group("/api/v1")
	{
		v1.GET("/test", func(context *gin.Context) {
			context.Writer.WriteString("asd")
		})

		// User Methods

		userMethods := v1.Group("/user")
		{
			userMethods.POST("/register", router.UserController.CreateUserEndpoint)
			userMethods.GET("/all", router.UserController.FindAllEndpoin)
		}

	}

	//handler.POST

}

func MainHandler(context *gin.Context) string {
	return "asd"

}
