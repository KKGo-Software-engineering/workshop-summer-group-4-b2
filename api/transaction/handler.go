package transaction

import (
	"net/http"
	"strconv"

	"github.com/KKGo-Software-engineering/workshop-summer/api/errs"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

type Handler interface {
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	GetExpenses(c echo.Context) error
	GetSummary(c echo.Context) error
	GetBalance(c echo.Context) error
	UpdateExpense(c echo.Context) error
	DeleteExpense(c echo.Context) error
}

func NewHandler(service Service) Handler {
	return handler{
		service: service,
	}
}

func (h handler) GetAll(c echo.Context) error {
	filter, ok := c.Get("filter").(Filter)
	if !ok {
		filter = Filter{}
	}

	pagination, ok := c.Get("pagination").(Pagination)
	if !ok {
		pagination = Pagination{}
	}

	result, err := h.service.GetAll(filter, pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h handler) Create(c echo.Context) error {
	return nil
}

func (h handler) GetExpenses(c echo.Context) error {
	return nil
}

func (h handler) GetSummary(c echo.Context) error {
	spenderId, err := strconv.Atoi(c.QueryParam("spender_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, errs.Build(err))
	}

	txnType := c.QueryParam("txn_type")

	summary, err := h.service.GetSummary(spenderId, txnType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errs.Build(err))
	}

	return c.JSON(http.StatusOK, summary)
}

func (h handler) GetBalance(c echo.Context) error {
	return nil
}

func (h handler) UpdateExpense(c echo.Context) error {
	return nil
}

func (h handler) DeleteExpense(c echo.Context) error {
	return nil
}
