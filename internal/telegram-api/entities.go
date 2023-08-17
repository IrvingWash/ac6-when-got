package telegramapi

type TelegramUpdate struct {
	UpdateID int             `json:"update_id"`
	Message  TelegramMessage `json:"message"`
}

type TelegramMessage struct {
	MessageID int          `json:"message_id"`
	Chat      TelegramChat `json:"chat"`
}

type TelegramChat struct {
	ID int `json:"id"`
}

type TelegramSendMessagePayload struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
