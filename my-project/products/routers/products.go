package routers

import (
	"github.com/gorilla/mux"
	"github.com/aj/my-project/products/controllers"
)

func SetCustomersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/products", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return router
}
