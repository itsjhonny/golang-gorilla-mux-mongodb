package routes

import (
	"mux-mongo-api/controllers"

	"github.com/gorilla/mux"
)

func RoleRoute(router *mux.Router) {
	//router.HandleFunc("/roles", controllers.GetAllRole()).Methods("GET")
	//router.HandleFunc("/role/{roleId}", controllers.GetARole()).Methods("GET")
	router.HandleFunc("/role", controllers.CreateRole()).Methods("POST")
}
