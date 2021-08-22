package main

import (
	"dddsample/internal/helloworldserver"
	"dddsample/rpc/helloworld"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"
)

type HelloServer struct {}

func main() {
  mux := http.NewServeMux()
  // http.HandleFunc("/task", model.TasksHandler)
  // http.HandleFunc("/task/", model.TaskHandler)
  // http.HandleFunc("/", handleRoot)

  // twirpHandler := haberdasher.NewHaberdasherServer(&haberdasherserver.Server{}, nil)
  twirpHandler := helloworld.NewHelloWorldServer(&helloworldserver.HelloWorldServer{}, nil)
  mux.Handle(twirpHandler.PathPrefix(),twirpHandler)
  
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Defaulting to port %s", port)
  }
  
  log.Printf("Listening on port %s", port)
  if err := http.ListenAndServe(":"+port, mux); err != nil {
    log.Fatal(err)
  }
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
