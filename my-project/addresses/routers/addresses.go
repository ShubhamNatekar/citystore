package routers

import (
	"github.com/gorilla/mux"
	"github.com/aj/my-project/addresses/controllers"
)

func SetCustomersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/addresses", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/addresses", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/addresses/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return router
}
