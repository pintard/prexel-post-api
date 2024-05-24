package handler

import (
	"encoding/json"
	"net/http"
	"prexel-post-api/src/model"
	"prexel-post-api/src/repository"
	"prexel-post-api/src/service"
	"strconv"
	"strings"
	"time"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imagePath, err := service.UploadImage(file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := model.PrexelPost{
		UserId:     parseInt64(r.FormValue("user_id")),
		Code:       r.FormValue("code"),
		Title:      r.FormValue("title"),
		Tags:       strings.Split(r.FormValue("tags"), ","),
		ImagePath:  imagePath,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}

	id, err := repository.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func parseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
