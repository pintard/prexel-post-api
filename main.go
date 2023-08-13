package main

import (
	"net/http"
	"os"
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
	password = GetEnv("DB_PASSWORD", "password")
)

var log *Logger = GetLoggerInstance()

func main() {
	var err error = db.Connect(host, port, user, password, dbname)

	if err != nil {
		log.Error("Error connecting to the database: " + err.Error())
		os.Exit(1)
	}

	defer db.DB.Close()

	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/posts", CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts", GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{uuid:[0-9]+}", GetPostHandler).Methods("GET")

	serverPort := ":8080"
	log.Info("Server is starting and listening on port " + serverPort)
	http.ListenAndServe(serverPort, router)
}
