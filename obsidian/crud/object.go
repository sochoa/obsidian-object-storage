package crud

import (
	"github.com/gin-gonic/gin"
)

func postObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func putObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func getObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func deleteObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func SetupObjectRouter(router *gin.Engine) {
	router.POST("/:bucket/:path", postObject)
	router.PUT("/:bucket/:path", putObject)
	router.GET("/:bucket/:path", getObject)
	router.DELETE("/:bucket/:path", deleteObject)
}
