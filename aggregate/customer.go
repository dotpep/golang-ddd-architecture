// Package aggregate holds our aggrets that combines many entities into
// a full object
package aggregate

import (
	"errors"

	"github.com/dotpep/golang-ddd-architecture/entity"
	"github.com/dotpep/golang-ddd-architecture/valueobject"
	"github.com/google/uuid"
)

var (
	ErrorInvalidPerson = errors.New("A customer has to have a valid name")
)

// Customer aggregate that holds Person, Products and Transactions.
// Aggregate should not be accessible directly to grab the data,
// and is not accesible from outside.
type Customer struct {
	// Person is the root entity of customer
	// which means person.ID is the main identifier for customer
	// lowerCase or pascalCase means that their not accessible
	// from other domain, and outside.
	// Entity has pointers because they can change state,
	// and we want to recflect across the whole runtime
	// when everything something changes.
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// Factory Method/Pattern
// for encapsulate complex logic and for creating instances
// without knowing anything about the actual implementation details.

// New Customer is a Factory to create a new customer aggregate
// it will validate that the name is not empty.
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrorInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
