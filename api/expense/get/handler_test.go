package get

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock service implementing the Service interface
type MockService struct {
	mock.Mock
}

func (m *MockService) GetAll(filter Filter, paginate Paginate) ([]Expense, error) {
	args := m.Called(filter, paginate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Expense), args.Error(1)
}

func TestHandler_GetAll(t *testing.T) {
	// Arrange
	e := echo.New()

	mockService := new(MockService)
	h := NewHandler(mockService)

	tests := []struct {
		name           string
		queryParams    map[string]string
		expectedStatus int
		expectedBody   string
		setupMock      func()
	}{
		{
			name: "valid request with all parameters",
			queryParams: map[string]string{
				"date":        "2022-01-01",
				"amount":      "123.45",
				"category":    "test-category",
				"itemPerPage": "5",
				"page":        "2",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"amount":200, "date":null, "category":"category1", "imageUrl":"urlOne", "note":"note", "spenderId":"1"}]`,
			setupMock: func() {
				date, _ := time.Parse("2006-01-02", "2022-01-01")
				amount := float32(123.45)
				filter := Filter{
					Date:     &date,
					Amount:   &amount,
					Category: "test-category",
				}
				pagination := Paginate{
					ItemPerPage: 5,
					Page:        2,
				}
				mockService.On("GetAll", filter, pagination).Return([]Expense{{ID: 1, Date: nil, Amount: 200, Category: "category1", ImageUrl: "urlOne", Note: "note", SpenderId: "1"}}, nil)
			},
		},
		// {
		// 	name: "invalid date format",
		// 	queryParams: map[string]string{
		// 		"date": "invalid-date",
		// 	},
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedBody:   `"Invalid date format: parsing time "invalid-date" as "2006-01-02": cannot parse "invalid-date" as "2006"`,
		// 	setupMock:      func() {},
		// },
		// Add more test cases as needed
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			q := req.URL.Query()
			for k, v := range tt.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Assert
			if assert.NoError(t, h.GetAll(c)) {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.JSONEq(t, tt.expectedBody, rec.Body.String())
			}

			mockService.AssertExpectations(t)
		})
	}
}
