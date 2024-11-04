package aggregate

import (
	"errors"

	"github.com/dotpep/golang-ddd-architecture/entity"
	"github.com/google/uuid"
)

var (
	ErrorMissingValue = errors.New("missing important values")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

// NewProduct is factory of Aggregate
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrorMissingValue
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) SetID(id uuid.UUID) {
	if p.item == nil {
		p.item = &entity.Item{}
	}

	p.item.ID = id
}

func (p *Product) GetName() string {
	return p.item.Name
}

func (p *Product) SetName(name string) {
	if p.item == nil {
		p.item = &entity.Item{}
	}

	p.item.Name = name
}

func (p *Product) GetItem() *entity.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}
