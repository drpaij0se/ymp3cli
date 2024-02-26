package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type Respose struct {
	Message string
}

func DefaultRpc(port int) {
	err := client.Login("1211479421081485313")
	if err != nil {
		fmt.Println("No discord detected")
	}
	for {
		time.Sleep(time.Second * 1)

		url := "http://localhost:" + fmt.Sprintf("%d", port) + "/currentSong"
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalln(err)

		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)

		}

		r := string(body)
		song := "Listening " + r
		var res Respose
		json.Unmarshal([]byte(r), &res)
		err = client.SetActivity(client.Activity{
			State:      "🎵🖥️",
			Details:    song,
			LargeImage: "skull",
			LargeText:  "🙏",
			SmallImage: "https://cdn.discordapp.com/attachments/907631182240436305/1211480323817340938/640px-Kanye_West_at_the_2009_Tribeca_Film_Festival_28crop_229.png?ex=65ee59f9&is=65dbe4f9&hm=c9b9b1df4fd29814824c2b499aaec07d4ed2da1f213fe85fb6667f5a652ead24&",
			Buttons: []*client.Button{
				{
					Label: "GitHub",
					Url:   "https://github.com/drpaij0se/ymp3cli",
				},
			},
		})

		if err != nil {
			fmt.Println("Error in rpc")
		}
		fmt.Print("")

	}

}
