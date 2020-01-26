package controllers

import (
	"github.com/aj/my-project/shops/models"
)

type (
	// For Get - /customers
	ShopsResource struct {
		Data []models.Shops `json:"data"`
	}
	// For Post/Put - /customers
	ShopResource struct {
		Data models.Shops `json:"data"`
	}
)
