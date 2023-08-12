package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	. "prexel-post-api/data"
	. "prexel-post-api/model"
)

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
	Posts[post.UUID] = post

	fmt.Println("Post", post)
	fmt.Println("Posts", Posts)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
