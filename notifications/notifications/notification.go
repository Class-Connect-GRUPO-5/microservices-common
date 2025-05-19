package notification

import (
	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats/email_templates"
)

type Notification interface {
	Type() string
	Encode() ([]byte, error)
	Decode(data []byte) error
	AsEmail() (email_templates.Email, error)
	AsPush() (notification_formats.PushNotification, error)
}
