package customer

import (
	"errors"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/google/uuid"
)

var (
	ErrorCustomerNotFound       = errors.New("the customer not found in the repository")
	ErrorFailedToAddCustomer    = errors.New("failed to add the customer")
	ErrorFailedToUpdateCustomer = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
