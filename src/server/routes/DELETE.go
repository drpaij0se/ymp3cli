package routes

import (
	"github.com/drpaij0se/ymp3cli/src/server/controllers/DELETE"
	"github.com/labstack/echo"
)

func Delete(e *echo.Echo) {
	e.DELETE("/delete", DELETE.Delete)

}
