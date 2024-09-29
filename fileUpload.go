package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// file_upload_for_search handles the file upload and indexing.
func file_upload_for_search(r *gin.Engine) {
	r.POST("/upload", func(ctx *gin.Context) {
		// Retrieve the uploaded file
		file, err := ctx.FormFile("document")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request, your file was not uploaded"})
			return
		}

		// Open the uploaded file
		fileContent, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in opening your file"})
			return
		}
		defer fileContent.Close()

		// Read the file content
		content, err := io.ReadAll(fileContent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in reading your file"})
			return
		}

		// Save the uploaded file to the server
		err = ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save the file"})
			return
		}

		// Connect to Elasticsearch and index the document
		err = elastiConnect(ctx, file.Filename, content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload document", "error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Uploaded and indexed successfully"})
	})
}
