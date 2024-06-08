package services

import (
	"testing"

	"congdinh.com/crm/models"
)

func TestCustomerService_GetAll(t *testing.T) {
	customerService := NewCustomerService()

	customers := customerService.GetAll()

	if len(customers) != 5 {
		t.Errorf("Expected 5 customers, but got %d", len(customers))
	}
}

func TestCustomerService_GetById(t *testing.T) {
	customerService := NewCustomerService()

	customer := customerService.GetById(3)

	if customer == nil {
		t.Errorf("Expected customer with ID 3, but got nil")
	} else if customer.ID != 3 {
		t.Errorf("Expected customer with ID 3, but got ID %d", customer.ID)
	}
}

func TestCustomerService_Create(t *testing.T) {
	customerService := NewCustomerService()

	newCustomer := models.Customer{
		ID:        6,
		Name:      "New Customer",
		Role:      "Tester",
		Email:     "new@domain.com",
		Phone:     "987654321",
		Contacted: false,
	}

	success, err := customerService.Create(newCustomer)

	if err != nil {
		t.Errorf("Expected Create to return nil error, but got %s", err.Error())
	}

	if !success {
		t.Error("Expected Create to return true, but got false")
	}

	customers := customerService.GetAll()

	if len(customers) != 6 {
		t.Errorf("Expected 6 customers after Create, but got %d", len(customers))
	}
}

func TestCustomerService_Update(t *testing.T) {
	customerService := NewCustomerService()

	updatedCustomer := models.Customer{
		ID:        4,
		Name:      "Updated Customer",
		Role:      "Developer",
		Email:     "updated@domain.com",
		Phone:     "123456789",
		Contacted: true,
	}

	success, err := customerService.Update(4, updatedCustomer)

	if err != nil {
		t.Errorf("Expected Update to return nil error, but got %s", err.Error())
	}

	if !success {
		t.Error("Expected Update to return true, but got false")
	}

	customer := customerService.GetById(4)

	if customer == nil {
		t.Errorf("Expected customer with ID 4 after Update, but got nil")
	} else if customer.Name != "Updated Customer" {
		t.Errorf("Expected customer name to be 'Updated Customer', but got '%s'", customer.Name)
	}
}

func TestCustomerService_Delete(t *testing.T) {
	customerService := NewCustomerService()

	success := customerService.Delete(2)

	if !success {
		t.Error("Expected Delete to return true, but got false")
	}

	customers := customerService.GetAll()

	if len(customers) != 4 {
		t.Errorf("Expected 4 customers after Delete, but got %d", len(customers))
	}

	deletedCustomer := customerService.GetById(2)

	if deletedCustomer != nil {
		t.Errorf("Expected customer with ID 2 to be deleted, but got customer with ID %d", deletedCustomer.ID)
	}
}
