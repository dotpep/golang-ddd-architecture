// Package memory is a In-Memory implementation of Customer Repository
package memory

import (
	"fmt"
	"sync"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/dotpep/golang-ddd-architecture/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	// here we can not reach person.ID
	//mr.customers.person.ID?xxx
	//return aggregate.Customer{}, nil
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrorCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	// check if the map is initialized
	// if is not we do it
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	// Make sure customer is already in repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrorFailedToAddCustomer)
	}

	// if it doesn't we added
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrorUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
