package rabbitmq

import "github.com/streadway/amqp"

func CreateChannel() (*amqp.Channel, error) {
	rabbitConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}
	defer rabbitConn.Close()

	rabbitCh, err := rabbitConn.Channel()
	if err != nil {
		return nil, err
	}
	defer rabbitCh.Close()

	return rabbitCh, nil
}
