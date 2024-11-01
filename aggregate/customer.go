// Package aggregate holds our aggrets that combines many entities into
// a full object
package aggregate

import (
	"github.com/dotpep/golang-ddd-architecture/entity"
	"github.com/dotpep/golang-ddd-architecture/valueobject"
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
