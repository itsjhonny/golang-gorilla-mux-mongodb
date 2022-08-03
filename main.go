package main

import (
	"fmt"
	"log"
	"mux-mongo-api/configs"
	"mux-mongo-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func routers(router *mux.Router) {
	//routes

	fmt.Println("Seting routers")

	routes.UserRoute(router)
	routes.RoleRoute(router)
}

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//api cfgs
	cfg_api, err := configs.LoadAPIConfigs()

	if err != nil {
		log.Fatal(nil)
	}

	routers(router)

	log.Fatal(http.ListenAndServe(":"+cfg_api.Port, router))
}
