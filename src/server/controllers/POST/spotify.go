package POST

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/drpaij0se/ymp3cli/src/server/tools"
	"github.com/labstack/echo"
)

//this functions works for check any kind of error of the client

func Spotify(c echo.Context) (err error) {
	var stdout, stderr bytes.Buffer
	var inputUrl tools.Url
	c.Response().WriteHeader(http.StatusCreated)

	err = json.NewDecoder(c.Request().Body).Decode(&inputUrl)

	url := inputUrl.Url

	if tools.ErrControl(c, "spotify", url, tools.S) {

		switch runtime.GOOS {
		case "linux", "darwin":
			cmd := exec.Command("sh", "-c", "spotdl "+url)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err = cmd.Run(); err != nil {
				log.Println(err)
				stderr.Write([]byte(err.Error()))

			}

		case "windows":
			cmd := exec.Command("cmd", "/c", ("spotdl " + url))
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err = cmd.Run(); err != nil {
				log.Println(err)
				stderr.Write([]byte(err.Error()))

			}

		default:
			fmt.Println("Your OS is not supported")
		}
		executedOut := stdout.String() + stderr.String()
		output := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
		// NOTE: PONER ESTO QUE NO SE VEAN LAS ANSI LETTERS
		//out := noansi.NoAnsi(output)

		json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": output, "status": "success"})

		tools.MoveSong()
	}

	return

}
