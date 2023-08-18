package telegramapi

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int  `json:"message_id"`
	Chat      Chat `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

type SendMessagePayload struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
