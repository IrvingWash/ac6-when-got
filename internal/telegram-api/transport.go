package telegramapi

import (
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

func (t *transport) getMe() string {
	resp, err := http.Get(t.reqEnv.getMeRequestMetainfo().url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
