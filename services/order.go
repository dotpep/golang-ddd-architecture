package services

import (
	"log"

	"github.com/dotpep/golang-ddd-architecture/domain/customer"
	"github.com/dotpep/golang-ddd-architecture/domain/customer/memory"
	"github.com/google/uuid"
)

// Configuration Pattern
// Service Generator

// OrderConfiguration is an alias for a function
// signature which will accept the OrderService and return an error
// pointer to *OrderService in parameter because
// we want to modify the service based on the configuration
type OrderConfiguration func(os *OrderService) error

// OrderService have CustomerRepository because
// whenever somebody makes an order they are a customer
// we need to handle the customer aggregate so need CustomerRepository in OrderService
type OrderService struct {
	customers customer.CustomerRepository
}

// NewOrderService is factory function
// for OrderService,
// takes as parameter variable (variadic) amount of OrderConfiguration
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// Example:

//NewOrderService(
//	WithMemoryCustomerRepository(), //or WithMongoCustomerRepository(),
//	WithLogging("debug"),
//	WithTracing(),
//)

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the OrderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

// Business Logic of OrderService
func (o *OrderService) CreateOrder(customerID uuid.UUID, products []uuid.UUID) error {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	// TODO: Get each Product, Ouchie no ProductRepository
	log.Println(c)

	return nil
}
