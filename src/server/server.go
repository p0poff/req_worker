package server

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"params"
	// "errors"
	// "model"
	// "my_jwt"
)

func GetServer(param params.Init) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "error")
	})

	// http.Handle("/test", testHendler(param))
	
	// http.HandleFunc("/login", handlerWrapper(param, loginHandler))

	// http.HandleFunc("/check", handlerWrapper(param, checkHandler))
	err := http.ListenAndServe(param.Port, nil)
	return err
}