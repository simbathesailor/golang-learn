package cfg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func createNewServer() *echo.Echo {

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found page")
	}

	server := echo.New()

	// Middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  `{"uri": "${uri}", "method":"${method}","status":"${status}","header", "${header}", "error":"${error}"}` + "\n",
		Skipper: skipLogging(),
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())
	server.Use(middleware.Gzip())
	server.Use(middleware.Secure())
	server.Use(middleware.CORS())
	//server.Use(middleware.CSRF())

	server.GET("/", rootPingHandler)
	server.GET("/rtw", rtwEntryPingHandler)

	return server
}

// skip elb health check logging
func skipLogging() func(context echo.Context) bool {
	return func(context echo.Context) bool {
		return strings.ToLower(context.Request().UserAgent()) == strings.ToLower("ELB-HealthCheck/2.0")
	}
}

func runServer(server *echo.Echo, config *AppConfig) error {
	return server.Start(fmt.Sprintf(":%s", config.Server.Port))
}
