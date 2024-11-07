package services

import (
	"context"
	"log"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/dotpep/golang-ddd-architecture/domain/customer"

	memoryCustomer "github.com/dotpep/golang-ddd-architecture/domain/customer/memory"
	mongoCustomer "github.com/dotpep/golang-ddd-architecture/domain/customer/mongo"
	"github.com/dotpep/golang-ddd-architecture/domain/product"
	memoryProduct "github.com/dotpep/golang-ddd-architecture/domain/product/memory"
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
	products  product.ProductRepository

	//TODO: billing billing.Service (subservice)
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
//	WithMemoryProductRepository(),
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
	cr := memoryCustomer.NewMemoryCustomerRepository()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	//cr, err := mongoCustomer.New(ctx, connectionString)
	//if err != nil {
	//	return err
	//}
	//return WithCustomerRepository(cr)

	return func(os *OrderService) error {
		cr, err := mongoCustomer.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := memoryProduct.NewMemoryProductRepository()

		for _, product := range products {
			if err := pr.Add(product); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

// Business Logic of OrderService
func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product
	var products []aggregate.Product
	var totalPrice float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		totalPrice += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return totalPrice, nil
}
