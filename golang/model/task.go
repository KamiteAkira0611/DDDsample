package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

type Task struct {
	ID int64 `datastore:"-" goon:"id"`
	Body string 
}

func GetTask(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello, GAE/go GetTask")
	n := goon.NewGoon(r)
	tasks := []Task{}
	q := datastore.NewQuery("Task")

	_, err := n.GetAll(q, &tasks)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request. %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, GAE/go CreateTask")
	n := goon.NewGoon(r)

	body := r.FormValue("body")
	g := &Task{Body: body}

	if body == "" {
		fmt.Fprintf(w, "invalid request. body is nil")
		return
	}

	key, err := n.Put(g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request. %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(key)
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, GAE/go DeleteTask")
	s := r.URL.Path[len("/task/delete/"):]
	id, iderr := strconv.ParseInt(s, 10, 64)
	if iderr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request. %s is not ID.\n", s)
		return
	}

	n := goon.NewGoon(r)
	task := &Task{ID: id}
	getErr := n.Get(task)
	if getErr != nil {
		fmt.Fprintf(w, "invalid request. getErr %s", getErr.Error())
		return
	}
	
	deleteErr := n.Delete(task)
	if deleteErr != nil {
		fmt.Fprintf(w, "invalid request. deleteErr %s", deleteErr.Error())
	}

	json.NewEncoder(w).Encode(task)
}


func TaskHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			fmt.Fprint(w, "GET hello!\n")
		case http.MethodPost:
			fmt.Fprint(w, "POST hello!\n")
		case http.MethodDelete:
			fmt.Fprint(w, "DELETE hello!\n")
			
		default:
			fmt.Fprint(w, "Method not allowed.\n")
	}
}