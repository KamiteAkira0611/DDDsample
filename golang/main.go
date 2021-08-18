package main

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
  "google.golang.org/appengine"
)

type Task struct {
  ID       string `datastore:"-" goon:"id"`
  Body     string `datastore:"Body,noindex"`
  Done     bool   `datastore:"Done,noindex"`
}

var datastoreClient *datastore.Client

func init()  {
  http.HandleFunc("/", handleRoot)
  
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

func handleRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, GAE/go")
  c := appengine.NewContext(r)
  fmt.Fprintf(w, "[AppID]: %v\n", appengine.AppID(c))
  fmt.Fprintf(w, "[Datacenter]: %v\n", appengine.Datacenter(c))
  fmt.Fprintf(w, "[DefaultVersionHostname]: %v\n", appengine.DefaultVersionHostname(c))
  fmt.Fprintf(w, "[InstanceID]: %v\n", appengine.InstanceID())
  fmt.Fprintf(w, "[IsAppEngine]: %v\n", appengine.IsAppEngine())
  fmt.Fprintf(w, "[IsDevAppServer]: %v\n", appengine.IsDevAppServer())
  fmt.Fprintf(w, "[IsFlex]: %v\n", appengine.IsFlex())
  fmt.Fprintf(w, "[IsStandard]: %v\n", appengine.IsStandard())
  fmt.Fprintf(w, "[ModuleName]: %v\n", appengine.ModuleName(c))
  fmt.Fprintf(w, "[RequestID]: %v\n", appengine.RequestID(c))
  fmt.Fprintf(w, "[ServerSoftware]: %v\n", appengine.ServerSoftware())
  serviceAccount, _ := appengine.ServiceAccount(c)
  fmt.Fprintf(w, "[ServiceAccount]: %v\n", serviceAccount)
  fmt.Fprintf(w, "[VersionID]: %v\n", appengine.VersionID(c))
}
