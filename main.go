package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the new search engine."})
	})

	file_upload_for_search(r)

	r.Run(":8081")
}
