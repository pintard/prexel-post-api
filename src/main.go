package main

import (
	"net/http"
	"os"
	"prexel-post-api/src/handler"
	"prexel-post-api/src/utils"

	"github.com/gorilla/mux"
)

var (
	host     = utils.GetEnv("DB_HOST", "localhost")
	port     = utils.GetEnv("DB_PORT", "5432")
	user     = utils.GetEnv("DB_USER", "prexel_user")
	password = utils.GetEnv("DB_PASSWORD", "password")
	dbname   = utils.GetEnv("DB_NAME", "prexeldb")
)

var log *utils.Logger = utils.GetLoggerInstance()

func main() {
	var err error = utils.Connect(host, port, user, password, dbname)

	if err != nil {
		log.Error("Error connecting to the database: " + err.Error())
		os.Exit(1)
	}

	defer utils.DB.Close()

	var router *mux.Router = mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/posts", handler.CreatePostHandler).Methods("POST")
	apiRouter.HandleFunc("/posts", handler.GetPostsHandler).Methods("GET")
	apiRouter.HandleFunc("/posts/{id:[0-9]+}", handler.GetPostHandler).Methods("GET")

	serverPort := ":8080"
	log.Info("Server is starting and listening on port " + serverPort)
	http.ListenAndServe(serverPort, router)
}
