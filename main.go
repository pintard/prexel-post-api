package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PrexelPost struct {
	UUID     int64  `json:"uuid"`
	Username string `json:"username" binding:"required"`
	Contact  string `json:"contact"`
	Content  string `json:"content" binding:"required"`
}

var posts = make(map[int64]PrexelPost)

func main() {
	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts", GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{uuid:[0-9]+}", GetPostHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post PrexelPost
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	err = json.Unmarshal(body, &post)

	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	post.UUID = rand.Int63()
	posts[post.UUID] = post

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := strconv.ParseInt(vars["uuid"], 10, 64)

	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	post, exists := posts[uuid]

	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
