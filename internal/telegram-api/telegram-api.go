package telegramapi

import "fmt"

const baseUrl = "https://api.telegram.org/bot"

var re = newRequestEnvironment(baseUrl)
var t = newTransport(&re)

func GetMe() {
	me := t.getMe()

	fmt.Println(me)
}
