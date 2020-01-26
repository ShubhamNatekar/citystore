package routers

import (
	"github.com/gorilla/mux"
	"github.com/aj/my-project/shops/controllers"
)

func SetCustomersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/shops", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/shops", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/shops/{id}", controllers.DeleteCustomer).Methods("DELETE")
	return router
}
