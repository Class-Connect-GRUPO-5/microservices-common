package notifications

import (
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_types"
)

type Notification interface {
	Type() string
	Encode() ([]byte, error)
	Decode(data []byte) error
	AsEmail() (notification_formats.Email, error)
	AsPush() (notification_formats.PushNotification, error)
}

func DecodeNotification(notificationType string, body []byte) (Notification, error) {
	var notification Notification
	notificationTypes := []Notification{
		&notification_types.WelcomeNotification{},
		&notification_types.InscriptionConfirmationNotification{},
		&notification_types.AuxTeacherAssignmentNotification{},
		&notification_types.NewTaskNotification{},
		&notification_types.TaskHandingConfirmationNotification{},
		&notification_types.TaskFeedbackNotification{},
		&notification_types.NewAnswerNotification{},
		&notification_types.NewForumCommentNotification{},
	}
	for _, notificationTypeInstance := range notificationTypes {
		if notificationTypeInstance.Type() == notificationType {
			notification = notificationTypeInstance
			break
		}
	}
	if notification == nil {
		return nil, fmt.Errorf("unknown notification type: %s", notificationType)
	}
	err := notification.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("error decoding notification: %v", err)
	}
	return notification, err
}
