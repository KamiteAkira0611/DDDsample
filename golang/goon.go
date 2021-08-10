package main

import (
	// "context"
	"fmt"
	"net/http"

  "github.com/mjibson/goon"
)

func getGoonTask(w http.ResponseWriter, r *http.Request) {
  u := &Task{ID: "123", Body: "aaa", Done: true}
  key, err := goon.NewGoon(r).Put(u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}

	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v", key)
}
