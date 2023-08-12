package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "prexel-post-api/model"

	"github.com/gorilla/mux"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	uuid, err := strconv.ParseInt(vars["uuid"], 10, 64)

	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	var post PrexelPost
	var exists bool
	println(uuid)
	// post, exists := Posts[uuid]

	if !exists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
