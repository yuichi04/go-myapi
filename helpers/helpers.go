package helpers

import (
	"github.com/labstack/echo/v4"
)

func ReturnErrorInJSON(c echo.Context, statusCode int, message string) error {
	errObj := map[string]string{"error": message}
	return c.JSON(statusCode, errObj)
}

func ReturnErrorInJSONPretty(c echo.Context, statusCode int, message string) error {
	errObj := map[string]string{"error": message}
	return c.JSONPretty(statusCode, errObj, "    ")
}
