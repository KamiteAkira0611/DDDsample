package main

import (
	"fmt"
	"net/http"
	"time"
	"log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Employee struct {
	Name     string
	Role     string
	HireDate time.Time
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	e1 := &Employee{
		Name:     "Joe Citizen",
		Role:     "Manager",
		HireDate: time.Now(),
	}
	
	log.Printf("aaaaaaaaaaaa")
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "employee", nil), &e1)
	log.Printf("aaaaaaaaaaaa")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("aaaaaaaaaaaa")
	
	log.Printf("aaaaaaaaaaaa")
	var e2 Employee
	if err = datastore.Get(ctx, key, &e2); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
	}

	fmt.Fprintf(w, "Stored and retrieved the Employee named %q", e2.Name)
}