package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/", index)
  r.GET("/task", getTask)
  r.DELETE("/task/:taskid", deleteTask)
  r.POST("/task", createTask)

  port = "3000"

  entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}

func index(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}

func getTask(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}

func deleteTask(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}

func createTask(gc *gin.Context) {
  gc.String(http.StatusOK, "hello gin!")
}
