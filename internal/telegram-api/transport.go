package telegramapi

import (
	"io"
	"log"
	"net/http"
)

type transport struct {
	requestsEnvironment *requestsEnvironment
}

func newTransport(re *requestsEnvironment) transport {
	return transport{
		requestsEnvironment: re,
	}
}

func (t *transport) getMe() string {
	resp, err := http.Get(t.requestsEnvironment.getMe().url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
