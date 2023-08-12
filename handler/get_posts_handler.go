package handler

import (
	"encoding/json"
	"net/http"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}
