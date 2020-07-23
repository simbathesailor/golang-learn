package cfg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func rootPingHandler(c echo.Context) error {
	return c.String(http.StatusOK, `rtw service is running under route route => /rtw-backend`)
}

func rtwEntryPingHandler(c echo.Context) error {
	return c.String(http.StatusOK, `rtw service is running`)
}
