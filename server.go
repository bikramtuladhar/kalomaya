package main

import (
	"example/controller"
	"example/service"

	"github.com/gin-gonic/gin"
)


var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main() {
	// https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1&ab_channel=PragmaticReviews
	server := gin.Default()
	server.GET("/videos", func(ctx *gin.Context){
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}