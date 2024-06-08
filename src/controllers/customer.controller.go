package controllers

import (
	"encoding/json"
	"net/http"

	"congdinh.com/crm/services"
	viewmodels "congdinh.com/crm/view-models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	ICustomerService services.ICustomerService
}

// NewCustomerController creates a new customer controller
func NewCustomerController(customerService services.ICustomerService) *CustomerController {
	return &CustomerController{
		ICustomerService: customerService,
	}
}

// RegisterRoutes registers the routes for the customer controller
func (cc *CustomerController) RegisterRoutes(router *mux.Router) {
	customers := router.PathPrefix("/api/v1/customers").Subrouter()

	customers.HandleFunc("", cc.GetCustomers).Methods("GET")
	customers.HandleFunc("/{id}", cc.GetCustomer).Methods("GET")
	customers.HandleFunc("", cc.CreateCustomer).Methods("POST")
	customers.HandleFunc("/{id}", cc.UpdateCustomer).Methods("PUT")
	customers.HandleFunc("/{id}", cc.DeleteCustomer).Methods("DELETE")
}

// GetCustomers godoc
// @Summary Show a list of customers
// @Description get customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} viewmodels.CustomerViewModel
// @Router /customers [get]
func (cc *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers := cc.ICustomerService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomer godoc
// @Summary Show a customer
// @Description get customer by ID
// @Tags customers
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Success 200 {object} viewmodels.CustomerViewModel
// @Router /customers/{id} [get]
func (cc *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the request and convert it to an integer
	id, err := uuid.Parse(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer := cc.ICustomerService.GetById(id)

	if customer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description add by json customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   customer  body viewmodels.CustomerCreateViewModel  true  "Add Customer"
// @Success 201  {object}  viewmodels.CustomerViewModel  "Successfully created"
// @Failure 400  {object}  nil  "Bad Request"
// @Router /customers [post]
func (cc *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer viewmodels.CustomerCreateViewModel

	// Decode the request body into newCustomer
	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new customer to the slice
	result, err := cc.ICustomerService.Create(newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result.ID == uuid.Nil {
		http.Error(w, "Failed to create the customer", http.StatusBadRequest)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusCreated) // HTTP 201
	json.NewEncoder(w).Encode(result)
}

// UpdateCustomer godoc
// @Summary Update an existing customer
// @Description update by json customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   id   path      string  true  "Customer ID"
// @Param   customer  body      viewmodels.CustomerEditViewModel  true  "Update Customer"
// @Success 200  {object}  viewmodels.CustomerViewModel  "Successfully updated"
// @Failure 400  {object}  nil  "Bad Request"
// @Router /customers/{id} [put]
func (cc *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the request and convert it to an uuid
	id, error := uuid.Parse(mux.Vars(r)["id"])

	if error != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var updatedCustomer viewmodels.CustomerEditViewModel

	// Decode the request body into updatedCustomer
	err := json.NewDecoder(r.Body).Decode(&updatedCustomer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the customer in the slice
	result, err := cc.ICustomerService.Update(id, updatedCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result.ID == uuid.Nil {
		http.Error(w, "Failed to update the customer", http.StatusBadRequest)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK) // HTTP 200
	json.NewEncoder(w).Encode(updatedCustomer)
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description delete by customer ID
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   id   path      string  true  "Customer ID"
// @Success 204  "Successfully deleted"
// @Failure 400  {object}  nil  "Bad Request"
// @Router /customers/{id} [delete]
func (cc *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the request and convert it to an uuid
	id, err := uuid.Parse(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	needDelete := cc.ICustomerService.GetById(id)

	if needDelete == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Delete the customer from the slice
	result := cc.ICustomerService.Delete(id)
	if !result {
		http.Error(w, "Failed to delete the customer", http.StatusBadRequest)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusNoContent) // HTTP 204
}
