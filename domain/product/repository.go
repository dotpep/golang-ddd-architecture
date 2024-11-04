package product

import (
	"errors"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/google/uuid"
)

var (
	ErrorProductNotFound       = errors.New("the product not found in the repository")
	ErrorFailedToAddProduct    = errors.New("failed to add the product")
	ErrorFailedToUpdateProduct = errors.New("failed to update the product")
	ErrorFailedToDeleteProduct = errors.New("failed to delete the product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
