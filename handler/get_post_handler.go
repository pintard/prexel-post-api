package handler

import (
	"encoding/json"
	"net/http"
	"prexel-post-api/db"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	uuid, err := strconv.ParseInt(vars["uuid"], 10, 64)

	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	post, err := db.GetPost(uuid)

	if err != nil {
		http.Error(w, "Failed to find Post by that UUID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
