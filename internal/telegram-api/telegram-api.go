package telegramapi

import (
	"fmt"
	"log"

	"github.com/IrvingWash/ac6-when-got/internal/utils"
)

var re = newRequestEnvironment(makeBaseUrl())
var t = newTransport(&re)

func GetMe() {
	me := t.getMe()

	fmt.Println(me)
}

func makeBaseUrl() string {
	var botToken, err = utils.ExtractEnvironmentVariable("BOT_TOKEN")

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s", botToken)
}
