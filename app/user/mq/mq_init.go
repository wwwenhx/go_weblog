package mq

import (
	"gin_gomicro/config"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitMq() (err error) {
	conn, err = amqp.Dial(config.MqAddress)
	if err != nil {
		return err
	}
	ch, err = conn.Channel()
	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		"userChannel",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewMQ() (*amqp.Connection, *amqp.Channel) {
	return conn, ch
}
