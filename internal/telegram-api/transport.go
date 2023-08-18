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

func (t *transport) sendMessage(payload *SendMessagePayload) Message {
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

	var message Message

	err = json.Unmarshal(body, &message)

	if err != nil {
		log.Fatal(err)
	}

	return message
}
