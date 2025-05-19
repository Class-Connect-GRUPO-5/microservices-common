package notification_types

import (
	"encoding/json"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats/email_templates"
)

type WelcomeNotification struct {
	Name string `json:"name"`
}

func (n *WelcomeNotification) Type() string {
	return "Welcome"
}

func (n *WelcomeNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *WelcomeNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *WelcomeNotification) AsEmail() (email_templates.Email, error) {
	return email_templates.WelcomeEmail(n.Name), nil
}

func (n *WelcomeNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "Welcome to Class Connect",
		Text:  "Hello " + n.Name + ", welcome to Class Connect!",
	}, nil
}
