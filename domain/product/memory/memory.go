package memory

import (
	"fmt"
	"sync"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/dotpep/golang-ddd-architecture/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mr.products {
		products = append(products, product)
	}
	// In this case, we will never return our error
	// but it's up to the interface to determine if it's possible
	// it's not up to the implementation it's up to the interface
	// so repository decides if we need to return an error or not
	return products, nil
}

func (mr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mr.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrorProductNotFound
}

func (mr *MemoryProductRepository) Add(newProduct aggregate.Product) error {
	mr.Lock()
	defer mr.Unlock()

	// check if the map is initialized
	// if is not we do it
	if mr.products == nil {
		mr.Lock()
		mr.products = make(map[uuid.UUID]aggregate.Product)
		mr.Unlock()
	}

	// Make sure product is already in repository
	if _, ok := mr.products[newProduct.GetID()]; ok {
		return fmt.Errorf("product already exists :%w", product.ErrorFailedToAddProduct)
	}

	// if it doesn't we added
	mr.products[newProduct.GetID()] = newProduct
	return nil
}

func (mr *MemoryProductRepository) Update(newProduct aggregate.Product) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[newProduct.GetID()]; ok {
		return fmt.Errorf("product does not exist: %w", product.ErrorFailedToUpdateProduct)
	}

	mr.products[newProduct.GetID()] = newProduct
	return nil
}

func (mr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[id]; ok {
		return fmt.Errorf("product does not exist: %w", product.ErrorFailedToDeleteProduct)
	}

	delete(mr.products, id)
	return nil
}
