package main

import (
    "github.com/gorilla/mux"
    "log"
    "ddd-sample/models"

    "cloud.google.com/go/datastore"
    "google.golang.org/appengine"
)


func main() {
    // ルーターのイニシャライズ
    r := mux.NewRouter()

    books = append(books, Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "Philip", LastName: "Williams"}})
    books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Johnson"}})

    // ルート(エンドポイント)
    r.HandleFunc("/api/tasks", models.GetTasks).Methods("GET")
    r.HandleFunc("/api/tasks/{id}", models.GetTask).Methods("GET")
    r.HandleFunc("/api/tasks", models.CreateTask).Methods("POST")
    r.HandleFunc("/api/tasks/{id}", models.UpdateTask).Methods("PUT")
    r.HandleFunc("/api/tasks/{id}", models.DeleteTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}
