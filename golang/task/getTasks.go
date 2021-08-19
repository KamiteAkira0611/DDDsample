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
	key, err := n.Put(g)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"key": "%s"}`, key)))
}