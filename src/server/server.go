package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"params"
	"auth"
	"errors"
)

type response struct {
	Code int `json:"code"`
	Result map[string]string `json:"result"`
}

func (resp response) get(code int, res map[string]string) (string, error) {
	resp.Code = code
	resp.Result = res
	slcB, err := json.Marshal(&resp)
	return string(slcB), err
}

func (resp response) getError(error string) string {
	_r, _ := resp.get(400, map[string]string{"error":error})
	return _r
}

func setHandler(w http.ResponseWriter, r *http.Request, param params.Init) {
	var resp response
	mqClient := getMqService(param)
	ref, err := mqClient.Send(r)
	if err != nil {
		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
	}

	strR, err := resp.get(200, map[string]string{"ref": ref})
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	fmt.Fprintf(w, strR)
}

func getHandler(w http.ResponseWriter, r *http.Request, param params.Init) {
	var resp response
	respClient := getRespService(param)
	workResponse, err := respClient.GetResp(r)
	if err != nil {
		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
	}

	strR, err := resp.get(200, map[string]string{"response": workResponse})
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	fmt.Fprintf(w, strR)
}

func handlerWrapper(param params.Init, f func(http.ResponseWriter, *http.Request, params.Init)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if !auth.IsAuth(r) {
			var resp response
			http.Error(w, resp.getError(errors.New("Bad authorization").Error()), http.StatusBadRequest)
 			return
		} 
		f(w, r, param)
	}
}

func GetServer(param params.Init) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "error")
	})

	http.HandleFunc("/set", handlerWrapper(param, setHandler))

	http.HandleFunc("/get", handlerWrapper(param, getHandler))

	err := http.ListenAndServe(param.Port, nil)
	return err
}