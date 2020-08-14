package params

import (
	"mq_client"
	"resp_client"
	"net/http"
)

type MqSrv interface {
	Send(*http.Request) (string, error)
	Init(string) error
}

type RespSrv interface {
	GetResp(*http.Request) (string, error)
}

type Init struct {
	Port string
	MqClient MqSrv
	RespClient RespSrv
}

func InitParams() (Init, error) {
	return InitManual()
}

func InitManual() (Init, error) {
	_mqClient := new(mq_client.RabbitMq)
	err := _mqClient.Init("amqp://guest:guest@localhost")
	return Init{
		Port: ":9000", 
		MqClient: _mqClient, 
		RespClient: new(resp_client.RespMock)}, err 
}