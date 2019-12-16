package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

var e *echo.Echo

func Build() {
	e = echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if e == nil {
		Build()
	}
	e.ServeHTTP(w, r)
}
