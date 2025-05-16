package notification

import "github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_types"

type Notification interface {
	Type() string
	Encode() ([]byte, error)
	Decode(data []byte) error
	AsEmail() (notification_types.Email, error)
	AsPush() (notification_types.PushNotification, error)
}
