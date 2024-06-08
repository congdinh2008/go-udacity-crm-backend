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
func (cs *CustomerService) GetAll() []models.Customer {
	return cs.Customers
}

// GetById method return a customer by ID
func (cs *CustomerService) GetById(id int) *models.Customer {
	for _, customer := range cs.Customers {
		if customer.ID == id {
			return &customer
		}
	}
	return nil
}

// Create method create a new customer
func (cs *CustomerService) Create(customer models.Customer) (bool, error) {
	// Check if the customer id already exists
	for _, c := range cs.Customers {
		if c.ID == customer.ID {
			return false, errors.New("customer with the same id already exists")
		}
	}

	cs.Customers = append(cs.Customers, customer)
	return true, nil
}

// Update method update a customer by ID
func (cs *CustomerService) Update(id int, customer models.Customer) (bool, error) {
	for i, c := range cs.Customers {
		if c.ID == id {
			cs.Customers[i] = customer
			return true, nil
		}
	}

	return false, errors.New("customer not found")
}

// Delete method delete a customer by ID
func (cs *CustomerService) Delete(id int) bool {
	for i, customer := range cs.Customers {
		if customer.ID == id {
			cs.Customers = append(cs.Customers[:i], cs.Customers[i+1:]...)
			return true
		}
	}
	return false
}
