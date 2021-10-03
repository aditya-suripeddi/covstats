package wrapper

import (
	"github.com/labstack/echo/v4"
)

type HttpSuccess struct {
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"Data source is mohfw"`
	Success bool        `json:"success" example:"true"`
}

type HttpFail struct {
	Code    int         `json:"code" example:"500"`
	Data    interface{} `json:"data" example:""`
	Message string      `json:"message" example:"Invalid request"`
	Success bool        `json:"success" example:"false"`
}

// Data returns wrapped success data
func Data(code int, data interface{}, message string, c echo.Context) error {
	props := &HttpSuccess{
		Code:    code,
		Data:    data,
		Message: message,
		Success: true,
	}
	indent4Spaces := "    "
	return c.JSONPretty(code, props, indent4Spaces)
}

func Error(code int, message string, c echo.Context) error {
	props := &HttpFail{
		Code:    code,
		Data:    nil,
		Message: message,
		Success: false,
	}
	indent4Spaces := "    "
	return c.JSONPretty(code, props, indent4Spaces)
}
