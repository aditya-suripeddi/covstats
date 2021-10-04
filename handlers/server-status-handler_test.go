package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// root route
func TestServerStatusHandler(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	successResponse := "{\n    \"code\": 200,\n    \"data\": \"Server is up and running\",\n    \"message\": \"Server has started\",\n    \"success\": true\n}\n"
	// Assertions
	if assert.NoError(t, ServerStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, successResponse, rec.Body.String())
	}

}
