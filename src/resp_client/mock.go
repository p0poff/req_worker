package resp_client

import (
	// "errors"
	"net/http"
)

type RespMock struct {
	Response string
}

func (resp *RespMock) GetResp(r *http.Request) (string, error)  {
	resp.Response = "mock response string"
	return resp.Response, nil
}