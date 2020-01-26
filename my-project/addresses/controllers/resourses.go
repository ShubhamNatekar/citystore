package controllers

import (
	"github.com/aj/my-project/addresses/models"
)

type (
	// For Get - /addresses
	AddressesResource struct {
		Data []models.Address `json:"data"`
	}
	// For Post/Put - /addresses
	AddressResource struct {
		Data models.Address `json:"data"`
	}
)
