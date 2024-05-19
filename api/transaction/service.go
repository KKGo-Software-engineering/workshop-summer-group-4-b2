package transaction

import (
	"errors"
)

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
	result, err := s.repository.Create(request)
	if err != nil {
		return CreateTransactionResponse{}, errors.New("can't create transaction")
	}

	return result, nil
}

func (s service) GetBalance(spenderId int) (BalanceResponse, error) {
	txn, err := s.repository.GetSummary(spenderId, []string{
		"income", "expense",
	})
	if err != nil {
		return BalanceResponse{}, errors.New("can't get balance")
	}

	totalAmountEarned := 0.0
	totalAmountSpended := 0.0
	for _, v := range txn {
		if v.TxnType == "income" {
			totalAmountEarned += v.Amount
		}
		if v.TxnType == "expense" {
			totalAmountSpended += v.Amount
		}
	}

	totalBalance := totalAmountEarned - totalAmountSpended
	return BalanceResponse{
		TotalAmountEarned: totalAmountEarned,
		TotalAmountSpend:  totalAmountSpended,
		TotalAmountSaved:  totalBalance,
	}, nil
}

func (s service) UpdateExpense(transaction Transaction) error {
	err := s.repository.UpdateExpense(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s service) DeleteExpense(id int) error {
	err := s.repository.DeleteExpense(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetSummary(spenderId int, txnType string) (SummaryResponse, error) {
	return SummaryResponse{}, nil
}

func (s service) GetExpenses(spenderId int) ([]Transaction, error) {
	return make([]Transaction, 0), nil
}
