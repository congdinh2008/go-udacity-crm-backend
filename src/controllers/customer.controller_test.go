package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"congdinh.com/crm/models"
	"congdinh.com/crm/services"
	viewmodels "congdinh.com/crm/view-models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func TestCustomerController_GetCustomers(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)
	// Create a new request
	req, err := http.NewRequest("GET", "/api/v1/customers", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new response recorder
	rr := httptest.NewRecorder()
	// Serve the request
	router := mux.NewRouter()
	customerController.RegisterRoutes(router)
	router.ServeHTTP(rr, req)
	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
	// Check the response body
	expectedResponse := []viewmodels.CustomerViewModel{}
	json.Unmarshal(rr.Body.Bytes(), &expectedResponse)
	actualResponse := customerService.GetAll()
	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Expected response body %v, but got %v", expectedResponse, actualResponse)
	}
}

func TestCustomerController_GetCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)

	// Create a new customer
	newCustomer := viewmodels.CustomerCreateViewModel{
		Name:      "Vinh Dinh",
		Role:      "Developer",
		Email:     "vinhdinh@example.com",
		Phone:     "123456789",
		Contacted: false,
	}
	// Add the new customer to the service
	result, err := customerService.Create(newCustomer)

	if err != nil {
		t.Errorf("Expected Create to return nil error, but got %s", err.Error())
	}

	if result.ID == uuid.Nil {
		t.Error("Expected Create to return a new customer ID, but got nil")
	}

	// Create a new request
	req, err := http.NewRequest("GET", "/api/v1/customers/"+result.ID.String(), nil)

	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	router := mux.NewRouter()
	customerController.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedResponse := result
	actualResponse := viewmodels.CustomerViewModel{}
	json.Unmarshal(rr.Body.Bytes(), &actualResponse)

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Expected response body %v, but got %v", expectedResponse, actualResponse)
	}

	// Check if the customer was created
	createdCustomer := customerService.GetById(result.ID)
	if createdCustomer == nil {
		t.Errorf("Expected customer with ID %s to be created, but got nil", result.ID.String())
	} else {
		if createdCustomer.Name != newCustomer.Name {
			t.Errorf("Expected customer with ID %s to have name %s, but got '%s'", result.ID.String(), newCustomer.Name, createdCustomer.Name)
		}
		if createdCustomer.Role != newCustomer.Role {
			t.Errorf("Expected customer with ID %s to have role %s, but got '%s'", result.ID.String(), newCustomer.Role, createdCustomer.Role)
		}
		if createdCustomer.Email != newCustomer.Email {
			t.Errorf("Expected customer with ID %s to have email %s, but got '%s'", result.ID.String(), newCustomer.Email, createdCustomer.Email)
		}
	}
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)

	// Create a new customer
	newCustomer := viewmodels.CustomerCreateViewModel{
		Name:      "Vinh Dinh",
		Role:      "Developer",
		Email:     "vinhdinh@example.com",
		Phone:     "123456789",
		Contacted: false,
	}
	// Create a new request body with the new customer
	reqBody, _ := json.Marshal(newCustomer)
	// Create a new request
	req, err := http.NewRequest("POST", "/api/v1/customers", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	router := mux.NewRouter()
	customerController.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rr.Code)
	}
	// Check the response body
	actualResponse := viewmodels.CustomerViewModel{}

	json.Unmarshal(rr.Body.Bytes(), &actualResponse)

	if actualResponse.ID == uuid.Nil {
		t.Error("Expected Create to return a new customer ID, but got nil")
	}

	// Check if the customer was created
	createdCustomer := customerService.GetById(actualResponse.ID)
	if createdCustomer == nil {
		t.Errorf("Expected customer with ID %s to be created, but got nil", actualResponse.ID.String())
	} else {
		if createdCustomer.Name != newCustomer.Name {
			t.Errorf("Expected customer with ID %s to have name %s, but got '%s'", actualResponse.ID.String(), newCustomer.Name, createdCustomer.Name)
		}
		if createdCustomer.Role != newCustomer.Role {
			t.Errorf("Expected customer with ID %s to have role %s, but got '%s'", actualResponse.ID.String(), newCustomer.Role, createdCustomer.Role)
		}
		if createdCustomer.Email != newCustomer.Email {
			t.Errorf("Expected customer with ID %s to have email %s, but got '%s'", actualResponse.ID.String(), newCustomer.Email, createdCustomer.Email)
		}
	}
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)

	// Create a new customer
	newCustomer := viewmodels.CustomerCreateViewModel{
		Name:      "Vinh Dinh",
		Role:      "Tester",
		Email:     "vinhdinh@example.com",
		Phone:     "987654321",
		Contacted: false,
	}
	// Add the new customer to the service
	result, err := customerService.Create(newCustomer)

	if err != nil {
		t.Errorf("Expected Create to return nil error, but got %s", err.Error())
	}

	if result.ID == uuid.Nil {
		t.Error("Expected Create to return a new customer ID, but got nil")
	}

	// Create a new request body with the updated customer
	updatedCustomer := models.Customer{
		ID:        result.ID,
		Name:      "Dinh Van Vinh",
		Role:      "Product Owner",
		Email:     "vinhdinh@example.com",
		Phone:     "0987654321",
		Contacted: true,
	}
	reqBody, _ := json.Marshal(updatedCustomer)

	// Create a new request
	req, err := http.NewRequest("PUT", "/api/v1/customers/"+updatedCustomer.ID.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	router := mux.NewRouter()
	customerController.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check if the customer was updated
	updatedCustomerEntity := *customerService.GetById(updatedCustomer.ID)

	if updatedCustomer.Name != updatedCustomerEntity.Name {
		t.Errorf("Expected customer with ID %s to have name '%s', but got '%s'", updatedCustomer.ID.String(), updatedCustomer.Name, updatedCustomerEntity.Name)
	}
	if updatedCustomer.Role != updatedCustomerEntity.Role {
		t.Errorf("Expected customer with ID %s to have role '%s', but got '%s'", updatedCustomer.ID.String(), updatedCustomer.Role, updatedCustomerEntity.Role)
	}
	if updatedCustomer.Email != updatedCustomerEntity.Email {
		t.Errorf("Expected customer with ID %s to have email '%s', but got '%s'", updatedCustomer.ID.String(), updatedCustomer.Email, updatedCustomerEntity.Email)
	}
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()

	// Create a new customer controller
	customerController := NewCustomerController(customerService)

	existingCustomerId, err := uuid.Parse("4405071c-2adc-499d-966f-3cfdfa1deedc")

	if err != nil {
		t.Errorf("Failed to parse existing customer ID: %s", err.Error())
	}

	// Create a new request
	req, err := http.NewRequest("DELETE", "/api/v1/customers/"+existingCustomerId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Serve the request
	router := mux.NewRouter()
	customerController.RegisterRoutes(router)
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, rr.Code)
	}

	// Check if the customer was deleted
	deletedCustomer := customerService.GetById(existingCustomerId)
	if deletedCustomer != nil {
		t.Errorf("Expected customer with ID %s to be deleted, but got customer with ID %d", existingCustomerId.String(), deletedCustomer.ID)
	}
}
