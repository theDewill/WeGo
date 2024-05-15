package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func handleFirst_get(w http.ResponseWriter, r *http.Request) {
	print(w.Header())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {

	router := mux.NewRouter()

	epoint_obj := api.epoints("hello")

	posts = append(posts, Post{ID: "1", Title: "thsi is title", Body: "this is body"})
	router.HandleFunc("/posts", handleFirst_get).Methods("GET")

	http.ListenAndServe(":8080", router)

}
