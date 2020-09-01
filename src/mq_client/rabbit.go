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
	Channel *my_amqp.Channel
	pullQueues map[string]my_amqp.Queue
	Queue string `json:"queue"`
	Message string `json:"message"`
	RbMsg rabbitMqMsg
}

type rabbitMqMsg struct {
	Ref string `json:"ref"`
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
	mq.RbMsg.Ref = ref
	mq.RbMsg.Message = mq.Message
	queue, err := mq.getQueue(mq.Queue)
	if err != nil {
		return "", err
	}
	message, err := json.Marshal(mq.RbMsg)
	if err != nil {
		return "", err
	}
	err = mq.Channel.Publish("", queue.Name, false, false, my_amqp.Publishing{
        DeliveryMode: my_amqp.Persistent,
        ContentType:  "text/plain",
        Body:         []byte(message),
    })
	return ref, err
}

func (mq *RabbitMq) getQueue(name string) (my_amqp.Queue, error) {
	queue, prs := mq.pullQueues[name]
	if prs == true {
		return queue, nil
	} 
	queue, err := mq.Channel.QueueDeclare(name, true, false, false, false, nil)
	if err == nil {
		mq.pullQueues[name] = queue
	}
	return queue, err
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
	mq.Ref = new(RndToken)
	mq.pullQueues = make(map[string]my_amqp.Queue)
	conn, err := my_amqp.Dial(url)
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	} 
	mq.Connect = conn
	mq.Channel = channel
	return err
}