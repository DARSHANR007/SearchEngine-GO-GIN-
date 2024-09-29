package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
)

func elastiConnect(ctx *gin.Context, fileName string, content []byte) error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error connecting to Elasticsearch", "error": err.Error()})
		return err
	}

	// Fetch Elasticsearch info
	info, err := es.Info()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Elasticsearch info", "error": err.Error()})
		return err
	}
	defer info.Body.Close()

	fmt.Println(info)

	req := esapi.IndexRequest{
		Index:      "documents",
		DocumentID: fileName,
		Body:       bytes.NewReader(content),
		Refresh:    "true",
	}

	response, err := req.Do(ctx, es)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error indexing the document", "error": err.Error()})
		return err
	}
	defer response.Body.Close()

	if response.IsError() {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error with Elasticsearch: %v", response.Status())})
		return fmt.Errorf("error with Elasticsearch: %v", response.Status())
	}

	fmt.Println("Document indexed successfully in Elasticsearch")
	ctx.JSON(http.StatusOK, gin.H{"message": "Document indexed successfully"})
	return nil
}
