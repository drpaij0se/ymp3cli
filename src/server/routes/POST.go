package routes

import (
	"github.com/drpaij0se/ymp3cli/src/server/controllers/POST"
	"github.com/labstack/echo"
)

func Post(e *echo.Echo) {
	e.POST("/download", POST.Download)
	e.POST("/spotify", POST.Spotify)
	e.POST("/y", POST.Y)
	e.POST("/y/all", POST.PlayAllSongsReq)
}
