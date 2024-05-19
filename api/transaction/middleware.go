package transaction

import (
	"github.com/labstack/echo/v4"
)

type middleware struct {
	middlewareService MiddlewareService
}

type Middleware interface {
	SetFilterExpense(next echo.HandlerFunc) echo.HandlerFunc
	SetPagination(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(middlewareService MiddlewareService) Middleware {
	return middleware{
		middlewareService: middlewareService,
	}
}

func (m middleware) SetFilterExpense(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := c.QueryParams()
		if len(queryParams) == 0 {
			c.Set("filter", nil)
		}

		result := m.middlewareService.SetFilter(queryParams)
		c.Set("filter", result)

		return next(c)
	}
}

func (m middleware) SetPagination(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := c.QueryParams()
		if len(queryParams) == 0 {
			c.Set("pagination", nil)
		}

		result := m.middlewareService.SetPagination(queryParams)
		c.Set("pagination", result)

		return next(c)
	}
}
