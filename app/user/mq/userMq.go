package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewUserMq() *MQ {
	conn, ch := NewMQ()
	if conn == nil {
		fmt.Println("conn is nil")
	}
	if ch == nil {
		fmt.Println("ch is nil")
	}
	mq := &MQ{
		conn: conn,
		ch:   ch,
	}
	return mq
}

func (mq *MQ) Publish(queueName string, body []byte) error {
	_, err := mq.ch.QueueDeclare(
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

	err = mq.ch.Publish(
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
