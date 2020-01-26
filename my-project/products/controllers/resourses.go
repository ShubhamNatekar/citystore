package controllers

import (
	"github.com/aj/my-project/products/models"
)

type (
	// For Get - /customers
	ProductsResource struct {
		Data []models.ProductTypes `json:"data"`
	}
	// For Post/Put - /customers
	ProductResource struct {
		Data models.ProductTypes `json:"data"`
	}
)
