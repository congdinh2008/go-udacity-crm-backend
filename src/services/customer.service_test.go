package services

import (
	"testing"

	viewmodels "congdinh.com/crm/view-models"
	"github.com/google/uuid"
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

	existingCustomerId, err := uuid.Parse("4405071c-2adc-499d-966f-3cfdfa1deedc")

	if err != nil {
		t.Errorf("Failed to parse existing customer ID: %s", err.Error())
	}

	customer := customerService.GetById(existingCustomerId)

	if customer == nil {
		t.Errorf("Expected customer with ID 3, but got nil")
	} else if customer.ID != existingCustomerId {
		t.Errorf("Expected customer with ID 3, but got ID %s", customer.ID.String())
	}
}

func TestCustomerService_Create(t *testing.T) {
	customerService := NewCustomerService()

	newCustomer := viewmodels.CustomerCreateViewModel{
		Name:      "New Customer",
		Role:      "Tester",
		Email:     "new@domain.com",
		Phone:     "987654321",
		Contacted: false,
	}

	result, err := customerService.Create(newCustomer)

	if err != nil {
		t.Errorf("Expected Create to return nil error, but got %s", err.Error())
	}

	if result.ID == uuid.Nil {
		t.Error("Expected Create to return a new customer ID, but got nil")
	}

	customers := customerService.GetAll()

	if len(customers) != 6 {
		t.Errorf("Expected 6 customers after Create, but got %d", len(customers))
	}

	newCustomerId := result.ID

	createdCustomer := customerService.GetById(newCustomerId)

	if createdCustomer == nil {
		t.Errorf("Expected customer with ID 6 after Create, but got nil")
	}

	if createdCustomer.Name != "New Customer" {
		t.Errorf("Expected customer name to be 'New Customer', but got '%s'", createdCustomer.Name)
	}

	if createdCustomer.Role != "Tester" {
		t.Errorf("Expected customer role to be 'Tester', but got '%s'", createdCustomer.Role)
	}

	if createdCustomer.Email != "new@domain.com" {
		t.Errorf("Expected customer email to be %s, but got %s", newCustomer.Email, createdCustomer.Email)
	}
}

func TestCustomerService_Update(t *testing.T) {
	customerService := NewCustomerService()

	existingCustomerId, err := uuid.Parse("4405071c-2adc-499d-966f-3cfdfa1deedc")

	if err != nil {
		t.Errorf("Failed to parse existing customer ID: %s", err.Error())
	}

	updatedCustomer := viewmodels.CustomerEditViewModel{
		ID:        existingCustomerId,
		Name:      "Updated Customer",
		Role:      "Developer",
		Email:     "updated@domain.com",
		Phone:     "123456789",
		Contacted: true,
	}

	result, err := customerService.Update(existingCustomerId, updatedCustomer)

	if err != nil {
		t.Errorf("Expected Update to return nil error, but got %s", err.Error())
	}

	if result.ID != existingCustomerId {
		t.Errorf("Expected Update to return customer with ID %s, but got ID %s", existingCustomerId.String(), result.ID.String())
	}

	updatedCustomerEntity := customerService.GetById(existingCustomerId)

	if updatedCustomerEntity == nil {
		t.Errorf("Expected customer with ID %s to be updated, but got nil", existingCustomerId.String())
	}

	if updatedCustomerEntity.Name != updatedCustomer.Name {
		t.Errorf("Expected customer name to be %s, but got '%s'", updatedCustomer.Name, updatedCustomerEntity.Name)
	}

	if updatedCustomerEntity.Role != updatedCustomer.Role {
		t.Errorf("Expected customer role to be %s, but got '%s'", updatedCustomer.Role, updatedCustomerEntity.Role)
	}

	if updatedCustomerEntity.Email != updatedCustomer.Email {
		t.Errorf("Expected customer email to be %s, but got %s", updatedCustomer.Email, updatedCustomerEntity.Email)
	}

	if updatedCustomerEntity.Phone != updatedCustomer.Phone {
		t.Errorf("Expected customer phone to be %s, but got %s", updatedCustomer.Phone, updatedCustomerEntity.Phone)
	}
}

func TestCustomerService_Delete(t *testing.T) {
	customerService := NewCustomerService()

	existingCustomerId, err := uuid.Parse("4405071c-2adc-499d-966f-3cfdfa1deedc")

	if err != nil {
		t.Errorf("Failed to parse existing customer ID: %s", err.Error())
	}

	success := customerService.Delete(existingCustomerId)

	if !success {
		t.Error("Expected Delete to return true, but got false")
	}

	customers := customerService.GetAll()

	if len(customers) != 4 {
		t.Errorf("Expected 4 customers after Delete, but got %d", len(customers))
	}

	deletedCustomer := customerService.GetById(existingCustomerId)

	if deletedCustomer != nil {
		t.Errorf("Expected customer with ID 2 to be deleted, but got customer with ID %d", deletedCustomer.ID)
	}
}
