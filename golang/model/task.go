package model

import (
	"fmt"
	"net/http"

	"github.com/mjibson/goon"
)

type Task struct {
	Body  string
}

func GetTask(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello, GAE/go GetTasks")
	
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, GAE/go CreateTasks")
	
	n := goon.NewGoon(r)
	g := &Task{Body: "name"}
	key, err := n.Put(g)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}

