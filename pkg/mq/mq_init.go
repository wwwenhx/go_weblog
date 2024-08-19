package mq

import (
	"fmt"
	"gin_gomicro/config"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Ch *amqp.Channel

func InitMq(ChannelName string) (err error) {
	fmt.Println(config.MqAddress)
	Conn, err = amqp.Dial(config.MqAddress)
	if err != nil {
		return err
	}
	Ch, err = Conn.Channel()
	if err != nil {
		return err
	}
	_, err = Ch.QueueDeclare(
		ChannelName,
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
	return Conn, Ch
}
