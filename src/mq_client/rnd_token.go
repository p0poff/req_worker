package mq_client

import (
  "crypto/rand"
  "fmt"
)

type RndToken struct {
	Token string
}

func (rt *RndToken) TokenGenerator() string {
	b := make([]byte, 6)
	rand.Read(b)
	rt.Token = fmt.Sprintf("%x", b)
	return rt.Token
}