package expense

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSetFilter(t *testing.T) {
	service := NewMiddlewareService()

	date := "2023-05-18"
	expectedDate, _ := time.Parse("2006-01-02", date)

	amount := "2000"
	expectedAmount, _ := strconv.ParseFloat(amount, 32)

	expectedCategory := "food"

	tests := []struct {
		test        string
		queryParams map[string][]string
		expected    Filter
	}{
		{
			test: "date is set in query params",
			queryParams: map[string][]string{
				"date": {date},
			},
			expected: Filter{
				Date: &expectedDate,
			},
		}, {
			test: "amount is set in query params",
			queryParams: map[string][]string{
				"amount": {amount},
			},
			expected: Filter{
				Amount: float32(expectedAmount),
			},
		}, {
			test: "category is set in query params",
			queryParams: map[string][]string{
				"category": {expectedCategory},
			},
			expected: Filter{
				Category: expectedCategory,
			},
		}, {
			test: "date and category is set in query params",
			queryParams: map[string][]string{
				"date":     {date},
				"category": {expectedCategory},
			},
			expected: Filter{
				Date:     &expectedDate,
				Category: expectedCategory,
			},
		}, {
			test: "all filters are set in query params",
			queryParams: map[string][]string{
				"date":     {date},
				"amount":   {amount},
				"category": {expectedCategory},
			},
			expected: Filter{
				Date:     &expectedDate,
				Amount:   float32(expectedAmount),
				Category: expectedCategory,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			filter := service.SetFilter(tt.queryParams)
			if !reflect.DeepEqual(filter, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, filter)
			}
		})
	}
}

func TestSetPagination(t *testing.T) {
	service := NewMiddlewareService()

	itemPerPage := "10"
	expectedItemPerPage, _ := strconv.Atoi(itemPerPage)

	page := "5"
	expectedPage, _ := strconv.Atoi(page)

	defaultItemPerPage := 5
	defaultPage := 2

	tests := []struct {
		test        string
		queryParams map[string][]string
		expected    Pagination
	}{
		{
			test: "only item per page is set in query params",
			queryParams: map[string][]string{
				"item-per-page": {itemPerPage},
			},
			expected: Pagination{
				ItemPerPage: expectedItemPerPage,
				Page:        defaultPage,
			},
		},
		{
			test: "only page is set in query params",
			queryParams: map[string][]string{
				"page": {page},
			},
			expected: Pagination{
				ItemPerPage: defaultItemPerPage,
				Page:        expectedPage,
			},
		},
		{
			test: "item per page and page are set in query params",
			queryParams: map[string][]string{
				"item-per-page": {itemPerPage},
				"page":          {page},
			},
			expected: Pagination{
				ItemPerPage: expectedItemPerPage,
				Page:        expectedPage,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			pagination := service.SetPagination(tt.queryParams)
			if !reflect.DeepEqual(pagination, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, pagination)
			}
		})
	}
}
