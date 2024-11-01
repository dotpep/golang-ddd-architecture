// Package memory is a In-Memory implementation of Customer Repository
package memory

import (
	"sync"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
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

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error){
	return aggregate.Customer{}, nil
}

func (mr *MemoryRepository) Add(customer aggregate.Customer) error {
	return nil
}

func (mr *MemoryRepository) Update(customer aggregate.Customer) error {
	return nil
}
