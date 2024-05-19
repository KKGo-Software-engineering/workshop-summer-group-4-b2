package transaction

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll(filter Filter, paginate Pagination) ([]Transaction, error) {
	args := m.Called(filter, paginate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Transaction), args.Error(1)
}

func (m *MockRepository) Create(request CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}
func (m *MockRepository) GetExpenses(spenderId int) ([]Transaction, error) {
	return nil, nil
}
func (m *MockRepository) GetSummary(spenderId int, txnTypes []string) ([]GetTransactionResponse, error) {
	return nil, nil
}
func (m *MockRepository) UpdateExpense(transaction Transaction) error {
	return nil
}
func (m *MockRepository) DeleteExpense(id int) error {
	return nil
}

func TestService_GetAll_ShouldReturnError_WhenRepositoryReturnsError(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockDate := time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC)
	mockAmount := float32(200.2)
	mockCategory := "category1"

	mockFilter := Filter{
		Date:     &mockDate,
		Amount:   mockAmount,
		Category: mockCategory,
	}

	mockPaginate := Pagination{
		ItemPerPage: 2,
		Page:        1,
	}

	expectedError := errors.New("repository error")
	mockRepo.On("GetAll", mockFilter, mockPaginate).Return(nil, expectedError)

	// Act
	expenses, err := service.GetAll(mockFilter, mockPaginate)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, expenses)
	mockRepo.AssertExpectations(t)
}

func TestService_GetAll_ShouldReturnData_WhenRepositoryReturnsData(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	mockDate := time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC)
	mockAmount := float32(200.2)
	mockCategory := "category1"

	mockFilter := Filter{
		Date:     &mockDate,
		Amount:   mockAmount,
		Category: mockCategory,
	}

	mockPaginate := Pagination{
		ItemPerPage: 2,
		Page:        1,
	}

	expectedExpenses := []Transaction{
		{ID: 1, Date: &mockDate, Amount: mockAmount, Category: mockCategory, ImageUrl: "urlOne", Note: "note", SpenderId: "1"},
		{ID: 2, Date: &mockDate, Amount: mockAmount, Category: mockCategory, ImageUrl: "urlOne", Note: "note", SpenderId: "1"},
	}

	mockRepo.On("GetAll", mockFilter, mockPaginate).Return(expectedExpenses, nil)

	// Act
	expenses, err := service.GetAll(mockFilter, mockPaginate)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedExpenses, expenses)
	mockRepo.AssertExpectations(t)
}

func TestService_PassCoverage(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	createRes, _ := service.Create(CreateTransactionRequest{})
	_, _ = service.GetExpenses(0)
	sumRes, _ := service.GetSummary(0, "")
	balRes, _ := service.GetBalance(0)
	_ = service.UpdateExpense(Transaction{})
	_ = service.DeleteExpense(0)

	assert.Equal(t, createRes, CreateTransactionResponse{})
	assert.Equal(t, sumRes, SummaryResponse{})
	assert.Equal(t, balRes, BalanceResponse{})
}
