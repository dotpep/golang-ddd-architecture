// Package memory is a In-Memory implementation of Customer Repository
package memory

import (
	"fmt"
	"sync"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/dotpep/golang-ddd-architecture/domain/customer"
	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	// here we can not reach person.ID
	//mr.customers.person.ID?xxx
	//return aggregate.Customer{}, nil
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrorCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(newCustomer aggregate.Customer) error {
	// check if the map is initialized
	// if is not we do it
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	// Make sure customer is already in repository
	if _, ok := mr.customers[newCustomer.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrorFailedToAddCustomer)
	}

	// if it doesn't we added
	mr.Lock()
	mr.customers[newCustomer.GetID()] = newCustomer
	defer mr.Unlock()
	return nil
}

func (mr *MemoryCustomerRepository) Update(newCustomer aggregate.Customer) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.customers[newCustomer.GetID()]; ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrorFailedToUpdateCustomer)
	}

	mr.customers[newCustomer.GetID()] = newCustomer

	return nil
}
