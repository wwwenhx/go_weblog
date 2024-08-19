package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MQ struct {
	Conn    *amqp.Connection
	Ch      *amqp.Channel
	ReplyTo string
	Msgs    <-chan amqp.Delivery
}

func NewMq() *MQ {
	conn, ch := NewMQ()
	if conn == nil {
		fmt.Println("conn is nil")
	}
	if ch == nil {
		fmt.Println("ch is nil")
	}
	mq := &MQ{
		Conn: conn,
		Ch:   ch,
	}
	return mq
}

func (mq *MQ) Publish(queueName string, body []byte) error {
	_, err := mq.Ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = mq.Ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (mq *MQ) SyncPublish(queueName string, body []byte) error {
	_, err := mq.Ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = mq.Ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			ReplyTo:     queueName,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
