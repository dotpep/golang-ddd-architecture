// Package entity holds all the entities that are shared across sudomains.
package entity

import "github.com/google/uuid"

// Person is an entity that represents
// a person in all domains,
// that can be mutable and unique identifer.
type Person struct {
	// ID and the identifier of the entity
	// UperCase or PascalCase because it can be changed,
	// Mutable.
	ID   uuid.UUID
	Name string
	Age  int
}
