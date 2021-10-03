package handlers

import (
	"net/http"

	"github.com/aditya-suripeddi/covstats/helpers/wrapper"
	"github.com/labstack/echo/v4"
)

// @Summary Check server status
// @Tags root
// @Accept application/json
// @Produce json
// @Success 200 {object}  wrapper.HttpSuccess{data=string}
// @Failure 500 {object}  wrapper.HttpFail
// @Router / [get]
func ServerStatus(c echo.Context) error {
	return wrapper.Data(http.StatusOK, "Server is up and running", "Server has started", c)
}
