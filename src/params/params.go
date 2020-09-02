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
	Init(resp_client.RespParam) error
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
	_respClient := new(resp_client.RespRedis)
	err = _respClient.Init(resp_client.RespParam{
		Addr: "localhost:6380",
		Password: "",
		DB: 0,
	})
	return Init{
		Port: ":9000", 
		MqClient: _mqClient, 
		RespClient: _respClient}, err 
}