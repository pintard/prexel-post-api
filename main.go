package main

import (
	"log"
	"net/http"

	"prexel-post-api/db"
	. "prexel-post-api/handler"
	. "prexel-post-api/utils"

	"github.com/gorilla/mux"
)

var (
	host     = GetEnv("DB_HOST", "localhost")
	port     = GetEnv("DB_PORT", "5432")
	dbname   = GetEnv("DB_NAME", "prexelpostdb")
	user     = GetEnv("DB_USER", "prexeluser")
	password = GetEnv("DB_PASSWORD", "your_password")
)

func main() {
	err := db.Connect(host, port, user, password, dbname)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	defer db.DB.Close()

	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts", GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{uuid:[0-9]+}", GetPostHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
