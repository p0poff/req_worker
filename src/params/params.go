package params

import (
	"mq_client"
	"resp_client"
	"net/http"
)

type MqSrv interface {
	Send(*http.Request) (string, error)
}

type RespSrv interface {
	GetResp(*http.Request) (string, error)
}

type Init struct {
	Port string
	MqClient MqSrv
	RespClient RespSrv
}

func InitParams() Init {
	return InitManual()
}

func InitManual() Init {
	return Init{
		Port: ":9000", 
		MqClient: new(mq_client.MqMock), 
		RespClient: new(resp_client.RespMock)} 
}