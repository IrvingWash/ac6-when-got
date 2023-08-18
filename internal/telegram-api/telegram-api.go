package telegramapi

import (
	"fmt"
	"log"

	"github.com/IrvingWash/ac6-when-got/internal/utils"
)

var reqEnv = newRequestsEnvironment(makeBaseUrl())
var transp = newTransport(&reqEnv)

func SendMessage(payload *SendMessagePayload) Message {
	return transp.sendMessage(payload)
}

func makeBaseUrl() string {
	var botToken, err = utils.ExtractEnvironmentVariable("BOT_TOKEN")

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("https://api.telegram.org/bot%s", botToken)
}
