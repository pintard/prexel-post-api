package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"prexel-post-api/db"
	"strconv"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var queryValues url.Values = r.URL.Query()

	lastUUIDStr, ok := queryValues["lastUUID"]
	if !ok || len(lastUUIDStr) < 1 {
		http.Error(w, "Missing 'lastUUID' parameter", http.StatusBadRequest)
		return
	}

	limitStr, ok := queryValues["limit"]
	if !ok || len(limitStr) < 1 {
		http.Error(w, "Missing 'limit' parameter", http.StatusBadRequest)
		return
	}

	lastUUID, err := strconv.ParseInt(lastUUIDStr[0], 10, 64)
	if err != nil {
		http.Error(w, "Invalid 'lastUUID' parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr[0])
	if err != nil {
		http.Error(w, "Invalid 'limit' parameter", http.StatusBadRequest)
		return
	}

	posts, err := db.PollPosts(lastUUID, limit)

	if err != nil {
		http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
