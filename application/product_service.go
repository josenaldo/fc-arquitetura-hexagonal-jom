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

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetAll() ([]ProductInterface, error) {
	products, err := s.Persistence.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
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
