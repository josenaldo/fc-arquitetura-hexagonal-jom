package dto

import "github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProductDto() *ProductDto {
	return &ProductDto{}
}

func (p *ProductDto) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, err
}
