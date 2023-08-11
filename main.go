package main

import (
	"net/http"

	. "prexel-post-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts", GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{uuid:[0-9]+}", GetPostHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
