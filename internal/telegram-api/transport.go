package telegramapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type transport struct {
	reqEnv *requestsEnvironment
}

func newTransport(reqEnv *requestsEnvironment) transport {
	return transport{
		reqEnv: reqEnv,
	}
}

func (t *transport) sendMessage(payload *TelegramSendMessagePayload) TelegramMessage {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(t.reqEnv.sendMessageURL(), "application/json", bytes.NewBuffer(payloadBytes))

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var message TelegramMessage

	json.Unmarshal(body, message)

	return message
}
