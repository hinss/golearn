package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()
	engine.GET("/ping/:name", func(ginCtx *gin.Context) {

		name := ginCtx.Param("name")

		ginCtx.JSON(200, gin.H{
			"message": name,
		})
	})

	engine.Run()

}
