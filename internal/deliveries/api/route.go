package api

import (
	"fmt"
	"github.com/dedekrnwan/go-clean/internal/config"
	"github.com/dedekrnwan/go-clean/internal/contract"
	"github.com/dedekrnwan/go-clean/internal/deliveries/api/handler/user"
	middlewares "github.com/dedekrnwan/go-clean/internal/deliveries/api/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *api) InitRoute(c *contract.Contract) {
	a.e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", config.Config.Server.AppName, config.Config.Server.Version)
		return c.String(http.StatusOK, message)
	})

	middlewares.Init(a.e)

	a.initV1Route(a.e.Group("/v1"), c)
}

func (a *api) initV1Route(g *echo.Group, c *contract.Contract) {
	user.NewHandler(c).Route(g.Group("/users"))
}
