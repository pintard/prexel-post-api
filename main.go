package main

import (
	"net/http"
	"os"
	"prexel-post-api/db"
	handler "prexel-post-api/handler"
	utils "prexel-post-api/utils"

	"github.com/gorilla/mux"
)

var (
	host     = utils.GetEnv("DB_HOST", "localhost")
	port     = utils.GetEnv("DB_PORT", "5432")
	user     = utils.GetEnv("DB_USER", "prexeluser")
	password = utils.GetEnv("DB_PASSWORD", "password")
	dbname   = utils.GetEnv("DB_NAME", "prexelpostdb")
)

var log *utils.Logger = utils.GetLoggerInstance()

func main() {
	var err error = db.Connect(host, port, user, password, dbname)

	if err != nil {
		log.Error("Error connecting to the database: " + err.Error())
		os.Exit(1)
	}

	defer db.DB.Close()

	var router *mux.Router = mux.NewRouter()
	router.HandleFunc("/posts", handler.CreatePostHandler).Methods("POST")
	router.HandleFunc("/posts", handler.GetPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", handler.GetPostHandler).Methods("GET")

	serverPort := ":8080"
	log.Info("Server is starting and listening on port " + serverPort)
	http.ListenAndServe(serverPort, router)
}
