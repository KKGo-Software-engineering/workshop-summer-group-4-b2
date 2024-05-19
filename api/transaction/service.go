package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetAll(filter Filter, pagination Pagination) ([]Transaction, error)
	Create(request CreateTransactionRequest) (CreateTransactionResponse, error)
	GetExpenses(spenderId int) ([]Transaction, error)
	GetSummary(spenderId int, txnType string) (SummaryResponse, error)
	GetBalance(spenderId int) (BalanceResponse, error)
	UpdateExpense(transaction Transaction) error
	DeleteExpense(id int) error
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

func (s service) Create(request CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}

func (s service) GetBalance(spenderId int) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}

func (s service) UpdateExpense(transaction Transaction) error {
	return nil
}

func (s service) DeleteExpense(id int) error {
	return nil
}

func (s service) GetSummary(spenderId int, txnType string) (SummaryResponse, error) {
	return SummaryResponse{}, nil
}

func (s service) GetExpenses(spenderId int) ([]Transaction, error) {
	return make([]Transaction, 0), nil
}