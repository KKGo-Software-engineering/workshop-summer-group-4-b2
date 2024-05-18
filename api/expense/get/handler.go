package get

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

type Handler interface {
	GetAll(c echo.Context) error
}

func NewHandler(service Service) Handler {
	return handler{
		service: service,
	}
}

func (h handler) GetAll(c echo.Context) error {

	filter := Filter{
		//Amount: float32(amount),
	}

	fmt.Println(filter)
	pagination := Paginate{}

	dateParam := c.QueryParam("date")
	if dateParam != "" {
		parsedDate, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid date format: %v", err))
		}
		filter.Date = &parsedDate
	}

	// Parse amount
	amountParam := c.QueryParam("amount")
	if amountParam != "" {
		amount, err := strconv.ParseFloat(amountParam, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid amount format: %v", err))
		}
		amount32 := float32(amount)
		filter.Amount = &amount32
	}

	// Parse category
	filter.Category = c.QueryParam("category")

	// Parse itemsPerPage
	itemPerPageParam := c.QueryParam("itemPerPage")
	if itemPerPageParam != "" {
		itemPerPage, err := strconv.Atoi(itemPerPageParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid itemsPerPage format: %v", err))
		}
		pagination.ItemPerPage = itemPerPage
	} else {
		// Set default itemsPerPage if not provided
		pagination.ItemPerPage = 10
	}

	// Parse page
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid page format: %v", err))
		}
		pagination.Page = page
	} else {
		// Set default page if not provided
		pagination.Page = 1
	}

	result, err := h.service.GetAll(filter, pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

