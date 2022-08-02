package main

import (
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//api cfgs
	cfg_api, err := configs.LoadAPIConfigs()

	if err != nil {
		log.Fatal(nil)
	}

	//routes
	routes.UserRoute(router)
	log.Fatal(http.ListenAndServe(":"+cfg_api.Port, router))
}
