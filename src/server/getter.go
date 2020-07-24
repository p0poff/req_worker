package server

import (
	"mq_client"
	"net/http"
)

type mq interface {
	Send(*http.Request) (string, error)
}

func getMqService() mq {
	return new(mq_client.MqMock)
}