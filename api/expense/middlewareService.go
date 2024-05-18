package expense

import (
	"strconv"
	"time"
)

type middlewareService struct{}

type MiddlewareService interface {
	SetFilter(queryParams map[string][]string) Filter
	SetPagination(queryParams map[string][]string) Pagination
}

func NewMiddlewareService() MiddlewareService {
	return middlewareService{}
}

func (m middlewareService) SetFilter(queryParams map[string][]string) Filter {
	filter := Filter{}
	for key, values := range queryParams {
		value := values[0]

		switch key {
		case "date":
			parseDate, err := time.Parse("2006-01-02", value)
			if err == nil {
				filter.Date = &parseDate
			}
		case "amount":
			amount, err := strconv.ParseFloat(value, 32)
			if err == nil {
				filter.Amount = float32(amount)
			}
		case "category":
			filter.Category = value
		}
	}

	return filter
}

func (m middlewareService) SetPagination(queryParams map[string][]string) Pagination {
	pagination := Pagination{
		ItemPerPage: 5,
		Page:        2,
	}

	for key, values := range queryParams {
		value := values[0]

		switch key {
		case "item-per-page":
			itemPerPage, err := strconv.Atoi(value)
			if err == nil {
				pagination.ItemPerPage = itemPerPage
			}
		case "page":
			page, err := strconv.Atoi(value)
			if err == nil {
				pagination.Page = page
			}
		}
	}

	return pagination
}
