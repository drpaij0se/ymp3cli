package routes

import (
	"github.com/drpaij0se/ymp3cli/src/server/controllers/GET"
	"github.com/labstack/echo"
)

func Get(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {

		return c.String(200, "Hello, World!")
	})

	e.GET("/songs", GET.Songs)

}
