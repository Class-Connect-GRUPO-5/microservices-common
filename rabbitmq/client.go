package rabbitmq

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Host string
	Port uint16
}

func Url(c Config) string {
	return fmt.Sprintf("amqp://guest:guest@%s:%d/", c.Host, c.Port)
}

type Client struct {
	name string
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewClient(name string, config Config, exchanges []string) (Client, error) {
	var conn *amqp.Connection
	for {
		var err error
		conn, err = amqp.Dial(Url(config))
		if err != nil {
		} else {
			break
		}
		time.Sleep(time.Second * 5)
	}
	ch, err := conn.Channel()
	if err != nil {
		return Client{}, fmt.Errorf("error getting rabbit channel: %s", err)
	}
	for _, exchangeName := range exchanges {

		err = ch.ExchangeDeclare(
			exchangeName,
			"fanout",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return Client{}, fmt.Errorf("failed to declare exchange: %v", err)
		}
	}
	return Client{name, conn, ch}, nil
}

func (r *Client) Send(exchange string, headers amqp.Table, body []byte) error {
	if r.ch == nil {
		return nil
	}
	headers["source"] = r.name
	return r.ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
