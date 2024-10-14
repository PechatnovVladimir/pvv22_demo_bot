package product

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {

	if idx > 0 && idx < len(allProducts) {
		return &allProducts[idx], nil
	}

	return nil, errors.New("некорректный код продукта")

}
