package handler

import (
	"encoding/json"
	"net/http"
	. "prexel-post-api/data"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Posts)
}
