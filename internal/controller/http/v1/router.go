package v1

import "github.com/gin-gonic/gin"

func NewRouter(handler *gin.Engine) {

	handler.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("as")
	})

}

func MainHandler(context *gin.Context) string {
	return "privet"
}
