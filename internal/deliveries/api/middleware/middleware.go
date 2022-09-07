package middlewares

import (
	"fmt"
	"github.com/dedekrnwan/go-clean/internal/config"
	"github.com/dedekrnwan/go-clean/pkg/utils/validator"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	var (
		APP  = config.Config.Server.AppName
		ENV  = os.Getenv("ENV")
		NAME = fmt.Sprintf("%s-%s", APP, ENV)
	)

	// e.Use(Context)
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		}),
		echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
			Format:           fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri}", NAME),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
}
