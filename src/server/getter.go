package server

import (
	"params"
)

func getMqService(param params.Init) params.MqSrv {
	return param.MqClient
}

func getRespService(param params.Init) params.RespSrv {
	return param.RespClient
}