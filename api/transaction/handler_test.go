package transaction

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock service implementing the Service interface
type MockService struct {
	mock.Mock
}

func (m *MockService) GetAll(filter Filter, paginate Pagination) ([]Transaction, error) {
	args := m.Called(filter, paginate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Transaction), args.Error(1)
}

func (m *MockService) Create(request CreateTransactionRequest) (CreateTransactionResponse, error) {
	return CreateTransactionResponse{}, nil
}
func (m *MockService) GetExpenses(spenderId int) ([]Transaction, error) {
	return nil, nil
}
func (m *MockService) GetSummary(spenderId int, txnType string) (SummaryResponse, error) {
	return SummaryResponse{}, nil
}
func (m *MockService) GetBalance(spenderId int) (BalanceResponse, error) {
	return BalanceResponse{}, nil
}
func (m *MockService) UpdateExpense(transaction Transaction) error {
	return nil
}
func (m *MockService) DeleteExpense(id int) error {
	return nil
}

func TestHandler_GetAll(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/expenses", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("filter", Filter{})
	c.Set("paginate", Pagination{})

	expected := []Transaction{
		{
			ID:        1,
			Amount:    2000.0,
			Date:      nil,
			Category:  "food",
			ImageUrl:  "http://img.png",
			Note:      "transaction note",
			SpenderId: 1,
		},
	}

	mockService := new(MockService)
	mockService.On("GetAll", mock.Anything, mock.Anything).Return(expected, nil).Once()
	h := NewHandler(mockService)

	err := h.GetAll(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestHandler_Create(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.Create(c)

	assert.NoError(t, err)
}

func TestHandler_GetExpenses(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions/expense/detail", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.GetExpenses(c)

	assert.NoError(t, err)
}

func TestHandler_GetSummary(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions/summary", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.GetSummary(c)

	assert.NoError(t, err)
}

func TestHandler_GetBalance(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions/balance", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.GetBalance(c)

	assert.NoError(t, err)
}

func TestHandler_UpdateExpense(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.UpdateExpense(c)

	assert.NoError(t, err)
}

func TestHandler_DeleteExpense(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/transactions/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := new(MockService)
	h := NewHandler(mockService)
	err := h.DeleteExpense(c)

	assert.NoError(t, err)
}
