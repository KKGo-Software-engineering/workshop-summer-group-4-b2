package api

import (
	"database/sql"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/KKGo-Software-engineering/workshop-summer/api/eslip"
	"github.com/KKGo-Software-engineering/workshop-summer/api/health"
	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/KKGo-Software-engineering/workshop-summer/api/spender"
	"github.com/KKGo-Software-engineering/workshop-summer/api/transaction"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	*echo.Echo
}

func New(db *sql.DB, cfg config.Config, logger *zap.Logger) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(mlog.Middleware(logger))

	v1 := e.Group("/api/v1")

	v1.GET("/slow", health.Slow)
	v1.GET("/health", health.Check(db))
	v1.POST("/upload", eslip.Upload)

	v1.Use(middleware.BasicAuth(AuthCheck))

	{
		middlewareService := transaction.NewMiddlewareService()
		middlewareHandler := transaction.NewMiddleware(middlewareService)

		repository := transaction.NewRepository(db)
		service := transaction.NewService(repository)
		handler := transaction.NewHandler(service)
		v1.GET("/transactions", handler.GetAll, middlewareHandler.SetFilterExpense, middlewareHandler.SetPagination)
		v1.POST("/transactions", handler.Create)
		v1.GET("/transactions/expense/detail", handler.GetExpenses)
		v1.GET("/transactions/summary", handler.GetSummary)
		v1.GET("/transactions/balance", handler.GetBalance)
		v1.PUT("/transactions/:id", handler.UpdateExpense)
		v1.DELETE("/transactions/:id", handler.DeleteExpense)
	}

	{
		h := spender.New(cfg.FeatureFlag, db)
		v1.GET("/spenders", h.GetAll)
		v1.POST("/spenders", h.Create)
	}

	return &Server{e}
}
