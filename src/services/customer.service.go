package services

// import Customer struct from models/customer.go
import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"runtime"

	"congdinh.com/crm/models"
	viewmodels "congdinh.com/crm/view-models"
	"github.com/google/uuid"
)

// CustomerService struct
type CustomerService struct {
	Customers []models.Customer
}

func readData() []models.Customer {
	// Open file
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../data/customers.json")
	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal JSON data into struct
	var customers []models.Customer
	err = json.Unmarshal(data, &customers)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return customers
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		Customers: readData(),
	}
}

// GetAl method return all customers
func (cs *CustomerService) GetAll() []viewmodels.CustomerViewModel {
	customerViewModels := []viewmodels.CustomerViewModel{}
	for _, customer := range cs.Customers {
		customerViewModel := viewmodels.CustomerViewModel{
			ID:        customer.ID,
			Name:      customer.Name,
			Role:      customer.Role,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Contacted: customer.Contacted,
		}
		customerViewModels = append(customerViewModels, customerViewModel)
	}
	return customerViewModels
}

// GetById method return a customer by ID
func (cs *CustomerService) GetById(id uuid.UUID) *viewmodels.CustomerViewModel {
	for _, customer := range cs.Customers {
		if customer.ID == id {
			customerViewModel := viewmodels.CustomerViewModel{
				ID:        customer.ID,
				Name:      customer.Name,
				Role:      customer.Role,
				Email:     customer.Email,
				Phone:     customer.Phone,
				Contacted: customer.Contacted,
			}
			return &customerViewModel
		}
	}
	return nil
}

// Create method create a new customer
func (cs *CustomerService) Create(customerCreateViewModel viewmodels.CustomerCreateViewModel) (viewmodels.CustomerViewModel, error) {
	// Check if the customer id already exists
	for _, c := range cs.Customers {
		if c.Email == customerCreateViewModel.Email || c.Phone == customerCreateViewModel.Phone {
			return viewmodels.CustomerViewModel{}, errors.New("customer already exists")
		}
	}

	newCustomer := models.Customer{
		ID:        uuid.New(),
		Name:      customerCreateViewModel.Name,
		Role:      customerCreateViewModel.Role,
		Email:     customerCreateViewModel.Email,
		Phone:     customerCreateViewModel.Phone,
		Contacted: customerCreateViewModel.Contacted,
	}

	cs.Customers = append(cs.Customers, newCustomer)

	customerViewModel := viewmodels.CustomerViewModel{
		ID:        newCustomer.ID,
		Name:      newCustomer.Name,
		Role:      newCustomer.Role,
		Email:     newCustomer.Email,
		Phone:     newCustomer.Phone,
		Contacted: newCustomer.Contacted,
	}

	return customerViewModel, nil
}

// Update method update a customer by ID
func (cs *CustomerService) Update(id uuid.UUID, customer viewmodels.CustomerEditViewModel) (viewmodels.CustomerViewModel, error) {
	var updatedCustomer models.Customer
	for i, c := range cs.Customers {
		if c.ID == id {
			updatedCustomer = models.Customer{
				ID:        id,
				Name:      customer.Name,
				Role:      customer.Role,
				Email:     customer.Email,
				Phone:     customer.Phone,
				Contacted: customer.Contacted,
			}
			cs.Customers[i] = updatedCustomer

			customerViewModel := viewmodels.CustomerViewModel{
				ID:        updatedCustomer.ID,
				Name:      updatedCustomer.Name,
				Role:      updatedCustomer.Role,
				Email:     updatedCustomer.Email,
				Phone:     updatedCustomer.Phone,
				Contacted: updatedCustomer.Contacted,
			}
			return customerViewModel, nil
		}
	}

	return viewmodels.CustomerViewModel{}, errors.New("customer not found")
}

// Delete method delete a customer by ID
func (cs *CustomerService) Delete(id uuid.UUID) bool {
	for i, customer := range cs.Customers {
		if customer.ID == id {
			cs.Customers = append(cs.Customers[:i], cs.Customers[i+1:]...)
			return true
		}
	}
	return false
}
