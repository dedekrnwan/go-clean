package rest

import (
	"fmt"
	"github.com/dedekrnwan/go-clean/config"
	"github.com/dedekrnwan/go-clean/internal/deliveries/rest/handler/user"
	middlewares "github.com/dedekrnwan/go-clean/internal/deliveries/rest/middleware"
	"github.com/dedekrnwan/go-clean/internal/factory"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *rest) InitRoute(f *factory.Factory) {
	r.e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", config.Config.Server.AppName, config.Config.Server.Version)
		return c.String(http.StatusOK, message)
	})

	middlewares.Init(r.e)

	r.initV1Route(r.e.Group("/v1"), f)
}

func (r *rest) initV1Route(g *echo.Group, f *factory.Factory) {
	user.NewHandler(f).Route(g.Group("/users"))
}
