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
	return fmt.Sprintf("%x", b)
}