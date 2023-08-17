package telegramapi

import (
	"log"
	"net/url"
)

type requestsEnvironment struct {
	baseUrl string
}

func newRequestsEnvironment(baseUrl string) requestsEnvironment {
	return requestsEnvironment{
		baseUrl: baseUrl,
	}
}

func (r *requestsEnvironment) sendMessageURL() string {
	sendMessageUrl, err := url.JoinPath(r.baseUrl, "sendMessage")

	if err != nil {
		log.Fatal(err)
	}

	return sendMessageUrl
}
