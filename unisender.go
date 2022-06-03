package unisendergo

import (
	"net/http"
	"time"
)

const base = "https://go1.unisender.ru/ru/transactional/api/v1"
const sendEmail = base + "/email/send.json"

type Unisender struct {
	apiKey string
	client *http.Client
}

func New(apiKey string, timeout time.Duration) *Unisender {
	client := http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost:     20,
			MaxIdleConnsPerHost: 2,
		},
		Timeout: timeout,
	}

	return &Unisender{
		client: &client,
		apiKey: apiKey,
	}
}
