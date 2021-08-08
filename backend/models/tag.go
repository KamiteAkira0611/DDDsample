package models

import (
	"net/http"
)

type Tag struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
}

func GetTags(w http.ResponseWriter, r *http.Request) {

}