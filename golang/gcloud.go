package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
)


func getDSTask(w http.ResponseWriter, r *http.Request)  {
  
}

func createDSTask(w http.ResponseWriter, r *http.Request)  {
  ctx := context.Background()
  var err error
  datastoreClient, err = datastore.NewClient(ctx, os.Getenv("DATASTORE_PROJECT_ID"))
  if err != nil {
    log.Fatal(err)
    return
  }

  u := &Task{ID: "123", Body: "aaa", Done: true}
  k := datastore.IncompleteKey("Task", nil)

  if _, err :=  datastoreClient.Put(ctx, k, u); err != nil {
    log.Fatal(err)
    return
  }

  fmt.Fprintf(w, "createDSTask 1 - OK, key: %+v", u)
}
