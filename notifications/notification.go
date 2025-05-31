package notifications

import (
	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

type Notification interface {
	Type() string
	Encode() ([]byte, error)
	Decode(data []byte) error
	AsEmail() (notification_formats.Email, error)
	AsPush() (notification_formats.PushNotification, error)
}
