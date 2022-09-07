package api

import (
	"context"
	"github.com/dedekrnwan/go-clean/internal/contract"
	"github.com/labstack/echo/v4"
)

type (
	Api interface {
		PrepareEcho() (func() error, func(ctx context.Context) error)
	}
	api struct {
		e *echo.Echo
		c *contract.Contract
	}
)

func New(c *contract.Contract) Api {
	e := echo.New()
	return &api{
		e,
		c,
	}
}

func (h *api) PrepareEcho() (func() error, func(ctx context.Context) error) {
	h.InitRoute(h.c)

	return func() error {
			return h.e.Start(":" + "8081")
		}, func(ctx context.Context) error {
			return h.e.Shutdown(ctx)
		}
}
