package main

import (
	telegramapi "github.com/IrvingWash/ac6-when-got/internal/telegram-api"
)

func main() {
	telegramapi.SendMessage(&telegramapi.TelegramSendMessagePayload{
		ChatID: 1234,
		Text:   "Hello",
	})
}
