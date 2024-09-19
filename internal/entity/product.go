package entity

import (
	"errors"
	"gordan/pkg/entity"
	"time"
)

var (
	ErrorIdIsRequired    = errors.New("id is required")
	ErrorIdIsInvalid     = errors.New("idd is invalid")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorPriceIsInvalid  = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorIdIsInvalid
	}
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorPriceIsInvalid
	}
	return nil
}
