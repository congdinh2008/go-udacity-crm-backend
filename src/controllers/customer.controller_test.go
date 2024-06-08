package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"congdinh.com/crm/models"
	"congdinh.com/crm/services"
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
	expectedResponse := []models.Customer{}
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
	newCustomer := models.Customer{
		ID:        6,
		Name:      "Vinh Dinh",
		Role:      "Developer",
		Email:     "vinhdinh@example.com",
		Phone:     "123456789",
		Contacted: false,
	}
	// Add the new customer to the service
	customerService.Create(newCustomer)
	// Create a new request
	req, err := http.NewRequest("GET", "/api/v1/customers/6", nil)
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
	// Check get customer by ID 6
	expectedResponse := customerService.GetById(6)
	actualResponse := newCustomer

	if !reflect.DeepEqual(actualResponse, *expectedResponse) {
		t.Errorf("Expected response body %v, but got %v", expectedResponse, actualResponse)
	}
}

func TestCustomerController_CreateCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)
	// Create a new customer
	newCustomer := models.Customer{
		ID:        6,
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
	// Check if the customer was created
	createdCustomer := customerService.GetById(newCustomer.ID)
	if createdCustomer == nil {
		t.Errorf("Expected customer with ID %d to be created, but got nil", newCustomer.ID)
	} else {
		if createdCustomer.Name != newCustomer.Name {
			t.Errorf("Expected customer with ID %d to have name %s, but got '%s'", newCustomer.ID, newCustomer.Name, createdCustomer.Name)
		}
		if createdCustomer.Role != newCustomer.Role {
			t.Errorf("Expected customer with ID %d to have role %s, but got '%s'", newCustomer.ID, newCustomer.Role, createdCustomer.Role)
		}
		if createdCustomer.Email != newCustomer.Email {
			t.Errorf("Expected customer with ID %d to have email %s, but got '%s'", newCustomer.ID, newCustomer.Email, createdCustomer.Email)
		}
	}
}

func TestCustomerController_UpdateCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()
	// Create a new customer controller
	customerController := NewCustomerController(customerService)
	// Create a new customer
	newCustomer := models.Customer{
		ID:        6,
		Name:      "Vinh Dinh",
		Role:      "Tester",
		Email:     "vinhdinh@example.com",
		Phone:     "987654321",
		Contacted: false,
	}
	// Add the new customer to the service
	customerService.Create(newCustomer)
	// Create a new request body with the updated customer
	updatedCustomer := models.Customer{
		ID:        6,
		Name:      "Dinh Van Vinh",
		Role:      "Product Owner",
		Email:     "vinhdinh@example.com",
		Phone:     "0987654321",
		Contacted: true,
	}
	reqBody, _ := json.Marshal(updatedCustomer)
	// Create a new request
	req, err := http.NewRequest("PUT", "/api/v1/customers/"+strconv.Itoa(updatedCustomer.ID), bytes.NewBuffer(reqBody))
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
	entity := *customerService.GetById(updatedCustomer.ID)
	if updatedCustomer.Name != entity.Name {
		t.Errorf("Expected customer with ID %d to have name '%s', but got '%s'", updatedCustomer.ID, updatedCustomer.Name, entity.Name)
	}
	if updatedCustomer.Role != entity.Role {
		t.Errorf("Expected customer with ID %d to have role '%s', but got '%s'", updatedCustomer.ID, updatedCustomer.Role, entity.Role)
	}
	if updatedCustomer.Email != entity.Email {
		t.Errorf("Expected customer with ID %d to have email '%s', but got '%s'", updatedCustomer.ID, updatedCustomer.Email, entity.Email)
	}
}

func TestCustomerController_DeleteCustomer(t *testing.T) {
	// Create a new customer service
	customerService := services.NewCustomerService()

	// Create a new customer controller
	customerController := NewCustomerController(customerService)

	customerId := 2

	// Create a new request
	req, err := http.NewRequest("DELETE", "/api/v1/customers/"+strconv.Itoa(customerId), nil)
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
	deletedCustomer := customerService.GetById(customerId)
	if deletedCustomer != nil {
		t.Errorf("Expected customer with ID %d to be deleted, but got customer with ID %d", customerId, deletedCustomer.ID)
	}
}
