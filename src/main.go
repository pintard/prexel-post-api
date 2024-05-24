package main

import (
	"net/http"
	"os"
	"prexel-post-api/src/handler"
	"prexel-post-api/src/utils"
	"prexel-post-api/src/utils/logger"

	"github.com/gorilla/mux"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		logger.Log.Error("Error loading configuration: " + err.Error())
		os.Exit(1)
	}

	if err := utils.InitDB(config); err != nil {
		logger.Log.Error("Error initializing the database: " + err.Error())
		os.Exit(1)
	}

	defer utils.CleanupDB(config)

	if err := utils.Connect(config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName); err != nil {
		logger.Log.Error("Error connecting to the database: " + err.Error())
		os.Exit(1)
	}

	defer utils.DB.Close()

	var router *mux.Router = mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/posts", handler.CreatePostHandler).Methods("POST")
	apiRouter.HandleFunc("/posts", handler.GetPostsHandler).Methods("GET")
	apiRouter.HandleFunc("/posts/{id:[0-9]+}", handler.GetPostHandler).Methods("GET")

	serverPort := ":8080"
	logger.Log.Info("Server is starting and listening on port " + serverPort)
	if err := http.ListenAndServe(serverPort, router); err != nil {
		logger.Log.Error("Server failed: " + err.Error())
	}
}
