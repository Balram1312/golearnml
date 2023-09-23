package main

import (
	"mlgo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.Routes(r)

	r.Run(":8000")

}
