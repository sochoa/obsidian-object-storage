package crud

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Serve(storageRoot string) {
	r := gin.Default()
	r.GET("/ping", ping)
	SetupObjectRouter(r, storageRoot)
	r.Run()
}
