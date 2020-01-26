package controllers

import (
	"github.com/aj/my-project/customers/models"
)

type (
	// For Get - /customers
	CustomersResource struct {
		Data []models.Customer `json:"data"`
	}
	// For Post/Put - /customers
	CustomerResource struct {
		Data models.Customer `json:"data"`
	}
)
