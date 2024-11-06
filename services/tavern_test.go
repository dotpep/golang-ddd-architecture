package services

import (
	"testing"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("ashley")
	if err != nil {
		t.Fatal(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(t)
	}

	//orders := []uuid.UUID{
	//	products[0].GetID(),
	//}

	orders := []uuid.UUID{}
	for _, prod := range products {
		orders = append(orders, prod.GetID())
	}

	// Execute Order
	err = tavern.Order(cust.GetID(), orders)

	if err != nil {
		t.Fatal(t)
	}
}
