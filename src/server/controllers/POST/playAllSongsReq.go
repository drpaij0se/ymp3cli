package POST

import (
	"encoding/json"

	"github.com/drpaij0se/ymp3cli/src/server/tools"
	"github.com/labstack/echo"
)

func PlayAllSongsReq(c echo.Context) error {
	var n tools.Nsong

	json.NewDecoder(c.Request().Body).Decode(&n)
	if n.Nsong == 1 {
		// play the song and send the response at the same time in  a goroutine

		tools.PlayAllSongs(c.Echo())

		json.NewEncoder(c.Response()).Encode(map[string]string{"shuffle": "Playing all songs"})
	}

	return nil
}
