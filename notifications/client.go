package notifications

import (
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

const NotificationsExchangeName = "notifications"

type notificationClient struct {
	rabbitmqClient rabbitmq.Client
}

var client *notificationClient

func Send(userId string, notification Notification) error {
	if client == nil {
		return fmt.Errorf("client not initialized")
	}
	body, err := notification.Encode()
	if err != nil {
		return fmt.Errorf("error encoding notification: %s", err)
	}
	return client.rabbitmqClient.Send(NotificationsExchangeName, amqp091.Table{"type": notification.Type(), "user": userId}, body)
}

type Config struct {
	ServiceName string
	Rabbitmq    rabbitmq.Config
}

func Init(config Config) error {
	rabbitmqClient, err := rabbitmq.NewClient(config.ServiceName, config.Rabbitmq, []string{NotificationsExchangeName})
	if err != nil {
		return fmt.Errorf("error connecting to rabbitmq: %s", err)
	}
	client = &notificationClient{
		rabbitmqClient: rabbitmqClient,
	}
	return nil
}
