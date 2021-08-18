package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
  // "google.golang.org/appengine"
)

type Task struct {
  ID       string `datastore:"-" goon:"id"`
  Body     string `datastore:"Body,noindex"`
  Done     bool   `datastore:"Done,noindex"`
}

var datastoreClient *datastore.Client

func init()  {
  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/goon/get_task", getGoonTask)
  
  http.HandleFunc("/gcloud/get_task", getDSTask)
  http.HandleFunc("/gcloud/create_task", createDSTask)
  
  http.HandleFunc("/appengine/handle", handle)
  http.HandleFunc("/appengine/get_task", getDSTask)
  http.HandleFunc("/appengine/create_task", createDSTask)
  port := "3000"

  log.Printf("Listening on port %s", port)
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal(err)
  }

}

func main() {
  log.Printf("start")
  // appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  fmt.Fprint(w, "Hello, World!")
}
