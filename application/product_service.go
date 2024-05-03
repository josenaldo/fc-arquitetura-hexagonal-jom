package application

import "errors"

type ProductService struct {
	Persistence ProductPersistenceInterface
}

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductPrice = errors.New("the price must be greater than zero")
	ErrRequiredProductName = errors.New("the name is required")
)

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, error := s.Persistence.Get(id)
	if error != nil {
		return nil, error
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	if price < 0 {
		return nil, ErrInvalidProductPrice
	}

	if name == "" {
		return nil, ErrRequiredProductName
	}

	product, err := NewProduct(name, price)
	if err != nil {
		return nil, err
	}

	productSaved, err := s.Persistence.Save(product)
	if err != nil {
		return nil, err
	}

	return productSaved, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}
