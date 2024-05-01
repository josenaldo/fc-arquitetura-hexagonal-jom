package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, error := s.Persistence.Get(id)
	if error != nil {
		return nil, error
	}

	return product, nil
}
