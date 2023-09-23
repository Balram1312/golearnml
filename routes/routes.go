package routes

import (
	"mlgo/handler"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", handler.FirstPage)
	r.GET("/health", handler.Health)
}
