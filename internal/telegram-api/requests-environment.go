package telegramapi

import (
	"log"
	"net/url"
)

type requestMetainfo struct {
	url string
}

type requestsEnvironment struct {
	baseUrl string
}

func newRequestEnvironment(baseUrl string) requestsEnvironment {
	return requestsEnvironment{
		baseUrl: baseUrl,
	}
}

func (r *requestsEnvironment) getMe() requestMetainfo {
	getMeUrl, err := url.JoinPath(r.baseUrl, "getMe")

	if err != nil {
		log.Fatal(err)
	}

	return requestMetainfo{
		url: getMeUrl,
	}
}
