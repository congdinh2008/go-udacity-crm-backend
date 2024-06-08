package services

import "congdinh.com/crm/models"

// ICustomerService defines the interface for customer service operations
type ICustomerService interface {
	GetAll() []models.Customer
	GetById(id int) *models.Customer
	Create(customer models.Customer) (bool, error)
	Update(id int, customer models.Customer) (bool, error)
	Delete(id int) bool
}
