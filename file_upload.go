package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func file_upload_for_search(r *gin.Engine) {
	r.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("document")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, your file was not uploaded"})
			return
		}


		ctx.JSON(http.StatusOK, gin.H{"message": "The file has been uploaded", "filename": file.Filename})

		err=ctx.SaveUploadedFile(file,"uploads/" + file.Filename)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save the file"})
			return
		}
	})
}
