package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetAll(filter Filter, pagination Pagination) ([]Transaction, error)
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) GetAll(filter Filter, paginate Pagination) ([]Transaction, error) {
	result, err := s.repository.GetAll(filter, paginate)
	if err != nil {
		return nil, err
	}

	return result, nil
}
