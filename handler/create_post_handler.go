package handler

import (
	"encoding/json"
	"io"
	"net/http"
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

	// TODO Call repository function
	post.Date = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
