package handler

import (
	"github.com/gin-gonic/gin"
)

func FirstPage(c *gin.Context) {

	ML()
	c.HTML(200, "index.html", gin.H{"title": "Code Studio"})
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"message": "connection is stable !!!"})
}
