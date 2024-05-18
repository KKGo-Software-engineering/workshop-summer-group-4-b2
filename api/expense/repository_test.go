package expense

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll_ShouldReturnError_WhenErrorOnPrepare(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewRepository(db)
	mock.ExpectPrepare(`SELECT id, date, amount, category, image_url, note, spender_id FROM transaction LIMIT \$1 OFFSET \$2`).WillReturnError(errors.New("error on prepare"))
	mockFilter := Filter{}

	mockPaginate := Pagination{
		ItemPerPage: 1,
		Page:        1,
	}
	// Act
	_, err = repo.GetAll(mockFilter, mockPaginate)

	// Assert
	assert.Error(t, err)
}

func TestGetAll_ShouldReturnError_WhenErrorOnScan(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewRepository(db)
	mock.ExpectPrepare(`SELECT id, date, amount, category, image_url, note, spender_id FROM transaction LIMIT \$1 OFFSET \$2`).ExpectQuery().WillReturnError(errors.New("error on scan"))
	mockFilter := Filter{}

	mockPaginate := Pagination{
		ItemPerPage: 1,
		Page:        1,
	}
	// Act
	_, err = repo.GetAll(mockFilter, mockPaginate)

	// Assert
	assert.Error(t, err)
}

func TestGetAll_ShouldReturnData_WhenCorrectInput(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error occurred while creating mock DB connection: %v", err)
	}

	repo := NewRepository(db)
	mockRows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "image_url", "note", "spender_id"}).
		AddRow("1", nil, "200.2", "category1", "urlOne", "note", "1").AddRow("2", nil, "400", "category2", "urlTwo", "note", "1")
	mock.ExpectPrepare(`SELECT id, date, amount, category, image_url, note, spender_id FROM transaction WHERE date = \$1 AND amount = \$2 AND category = \$3 LIMIT \$4 OFFSET \$5`).ExpectQuery().WillReturnRows(mockRows)

	mockDate := time.Date(2020, time.April,
		11, 21, 34, 01, 0, time.UTC)

	mockAmount := float32(10.0)
	mockCategory := "mock category"

	mockFilter := Filter{
		Date:     &mockDate,
		Amount:   mockAmount,
		Category: mockCategory,
	}

	mockPaginate := Pagination{
		ItemPerPage: 1,
		Page:        1,
	}

	expecteds := []Expense{
		{
			ID:        1,
			Date:      nil,
			Amount:    200.2,
			Category:  "category1",
			ImageUrl:  "urlOne",
			Note:      "note",
			SpenderId: "1",
		},
		{
			ID:        2,
			Date:      nil,
			Amount:    400,
			Category:  "category2",
			ImageUrl:  "urlTwo",
			Note:      "note",
			SpenderId: "1",
		},
	}
	// Act
	expenses, err := repo.GetAll(mockFilter, mockPaginate)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 2, len(expenses))

	for i, expected := range expecteds {
		assert.Equal(t, expected, expenses[i])
	}
}
