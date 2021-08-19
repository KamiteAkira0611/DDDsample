package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

type Task struct {
	ID int64 `datastore:"-" goon:"id"`
	Body  string
}

func GetTask(w http.ResponseWriter, r *http.Request)  {
	n := goon.NewGoon(r)
	list := []Task{}
	q := datastore.NewQuery("Task")

	_, err := n.GetAll(q, &list)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}
	json.NewEncoder(w).Encode(list)
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	n := goon.NewGoon(r)
	g := &Task{Body: "name"}

	key, err := n.Put(g)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"error": "Failed save config: %s"}`, err.Error())))
		return
	}
	json.NewEncoder(w).Encode(key)
}

