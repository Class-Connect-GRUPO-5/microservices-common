package notifications

import (
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/rabbitmq"
)

type notificationClient struct {
	rabbitmqClient rabbitmq.Client
}

var client *notificationClient

func Send() error {
	if client == nil {
		return fmt.Errorf("client not initialized")
	}
	return nil
}

type Config struct {
	ServiceName string
	Rabbitmq    rabbitmq.Config
}

func Init(config Config) error {
	rabbitmqClient, err := rabbitmq.NewClient(config.ServiceName, config.Rabbitmq, []string{"notifications"})
	if err != nil {
		return fmt.Errorf("error connecting to rabbitmq: %s", err)
	}
	client = &notificationClient{
		rabbitmqClient: rabbitmqClient,
	}
	return nil
}
