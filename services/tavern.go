package services

import (
	"log"

	"github.com/google/uuid"
)

// TavernConfiguration is used Configuration Pattern
type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	// OrederService to takes orders
	OrederService *OrderService

	// BillingService
	BillingService interface{}
}

// NewTavern is used Factory Pattern
func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrederService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrederService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer: %0.0f\n", price)

	return nil
}
