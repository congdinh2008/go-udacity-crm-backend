package services

import (
	viewmodels "congdinh.com/crm/view-models"
	"github.com/google/uuid"
)

// ICustomerService defines the interface for customer service operations
type ICustomerService interface {
	GetAll() []viewmodels.CustomerViewModel
	GetById(id uuid.UUID) *viewmodels.CustomerViewModel
	Create(customer viewmodels.CustomerCreateViewModel) (viewmodels.CustomerViewModel, error)
	Update(id uuid.UUID, customer viewmodels.CustomerEditViewModel) (viewmodels.CustomerViewModel, error)
	Delete(id uuid.UUID) bool
}
