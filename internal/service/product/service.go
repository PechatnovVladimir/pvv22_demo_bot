package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) GetByID(idx int) string {
	if idx >= 0 && idx < len(allProducts) {
		return allProducts[idx].Title
	}
	return "номер за пределами разумного"
}
