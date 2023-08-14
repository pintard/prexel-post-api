package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"prexel-post-api/db"
	. "prexel-post-api/model"
	"time"
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

	post.Date = time.Now()

	id, err := db.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
