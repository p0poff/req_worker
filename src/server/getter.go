package server

import (
	"mq_client"
	"resp_client"
	"net/http"
)

type mqSrv interface {
	Send(*http.Request) (string, error)
}

type respSrv interface {
	GetResp(*http.Request) (string, error)
}

func getMqService() mqSrv {
	return new(mq_client.MqMock)
}

func getRespService() respSrv {
	return new(resp_client.RespMock)
}