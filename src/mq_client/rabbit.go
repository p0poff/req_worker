package mq_client

import (
	"errors"
	// "fmt"
	"encoding/json"
	"net/http"
	my_amqp "github.com/streadway/amqp"
)

type rndTokenInterface interface {
	TokenGenerator() string
}

type RabbitMq struct {
	Ref rndTokenInterface
	Connect *my_amqp.Connection
	Queue string `json:"queue"`
	Message string `json:"message"`
}

func (mq *RabbitMq) Send(r *http.Request) (string, error)  {
	err := json.NewDecoder(r.Body).Decode(mq)
	if err != nil {
		return "", err
	}

	if err := mq.valid(); err != nil {
		return "", err
	}
	
	ref := mq.Ref.TokenGenerator()
	return ref, nil
}

func (mq *RabbitMq) valid() error {
	if mq.Queue == "" {
		return errors.New("Param Queue is wrong")
	}

	if mq.Message == "" {
		return errors.New("Param Message is wrong")
	}

	return nil
}

func (mq *RabbitMq) Init(url string) error {
	conn, err := my_amqp.Dial(url)
	if err == nil {
		mq.Connect = conn
	}
	return err
}