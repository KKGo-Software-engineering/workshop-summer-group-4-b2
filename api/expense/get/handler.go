package get

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
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
	//amount, err := strconv.ParseFloat(c.QueryParams("amount"), 32)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, err.Error())
	//}

	filter := Filter{
		//Amount: float32(amount),
	}

	fmt.Println(filter)
	pagination := Paginate{}

	result, err := h.service.GetAll(filter, pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
