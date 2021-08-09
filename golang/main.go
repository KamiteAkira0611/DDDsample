package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "context"
  "time"

  "cloud.google.com/go/datastore"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/", index)
  r.GET("/task", getTask)
  r.DELETE("/task/:taskid", deleteTaskk)
  r.POST("/task", createTask)

  port = "3000"

  entryPoint := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(entryPoint)
}

func index(gc *gin.Context)  {
  gc.String(http.StatusOK, "hello gin!")
}

func getTask(gc *gin.Context)  {
  ctx := context.Background()
}
