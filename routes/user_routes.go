package routes

import (
	"mux-mongo-api/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")

	/*
		router.HandleFunc("/user", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")

			json.NewEncoder(rw).Encode(map[string]string{"data": "user route"})
		}).Methods("GET")
	*/
}
