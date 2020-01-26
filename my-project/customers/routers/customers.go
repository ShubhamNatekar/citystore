package routers

import (
	"github.com/gorilla/mux"
	"github.com/aj/my-project/customers/controllers"
)

func SetCustomersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return router
}
