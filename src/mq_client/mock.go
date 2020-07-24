package mq_client

import (
	// "errors"
	"net/http"
)

type MqMock struct {
	Ref string
}

func (mq *MqMock) Send(r *http.Request) (string, error)  {
	mq.Ref = "mock ref string"
	return mq.Ref, nil
}