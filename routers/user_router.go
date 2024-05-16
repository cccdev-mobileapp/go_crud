package routers

import (
	"go-crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user", controllers.AddNewUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
}
