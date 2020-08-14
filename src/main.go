package main

import (
	"fmt"
	"params"
	"server"
)

func main() {
	fmt.Println("request worker go...")
	param, err := params.InitParams()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = server.GetServer(param)
	if err != nil {
        fmt.Println(err)
    }
}