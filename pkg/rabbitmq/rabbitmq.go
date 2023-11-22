package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume("order", "go-consumer", false, false, false, false, nil)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
