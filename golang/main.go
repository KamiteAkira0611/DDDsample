package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/", index)
  r.GET("/task", getTask)
  r.DELETE("/task/:taskid", deleteTask)
  r.POST("/task", createTask)

  port := "3000"

  log.Printf("Listening on port %s", port)
  entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}

func getDatastoreClient(ctx context.Context) (client *datastore.Client, err error) {
	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	client, err = datastore.NewClient(ctx, projectID)
	return
}

func index(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}

func getTask(gc *gin.Context) {
  ctx := context.Background()

  client, err := getDatastoreClient(ctx)
  if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
  log.Printf("client %s", client)
}

func deleteTask(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}

func createTask(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}
