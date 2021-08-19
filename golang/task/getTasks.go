package task

import (
	"fmt"
	"net/http"

	"github.com/mjibson/goon"
)

func GetTask(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello, GAE/go GetTasks")
	
	n := goon.NewGoon(r)
	g := &Task{Name: "name"}
	key, _ := n.Put(g)
	fmt.Fprintf(w, "<br>PUT 2 - OK, key: %+v", key)
}