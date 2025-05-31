package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
)

// NewTaskNotification represents a notification sent to users when a new task is assigned in a course.
type NewTaskNotification struct {
	CourseName  string `json:"course_name"`
	Title       string `json:"heading"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

func (n *NewTaskNotification) Type() string {
	return "NewTask"
}

func (n *NewTaskNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *NewTaskNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *NewTaskNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "New Task in " + n.CourseName,
		Text:  n.Title,
	}, nil
}

func (n *NewTaskNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "New Task: " + n.Title + " - " + n.CourseName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>New Task Assigned</title>
    <style>
        body {
            background-color: #f4f7fb;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 0;
        }

        .container {
            max-width: 520px;
            margin: 48px auto;
            background-color: #fff;
            border-radius: 18px;
            box-shadow: 0 8px 32px rgba(79, 70, 229, 0.08);
            overflow: hidden;
            border: 1px solid #e0e7ef;
        }

        .header {
            background: linear-gradient(90deg, #059669 0%%, #10b981 100%%);
            color: white;
            text-align: center;
            padding: 36px 24px 20px 24px;
            font-size: 28px;
            font-weight: 600;
            letter-spacing: 1px;
            border-top-left-radius: 18px;
            border-top-right-radius: 18px;
        }

        .content {
            padding: 32px 32px 24px 32px;
            text-align: left;
            font-size: 18px;
            color: #22223b;
        }

        .content p {
            margin: 18px 0;
        }

        .task-info {
            background-color: #f0fdf4;
            border-left: 4px solid #10b981;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .task-title {
            font-size: 20px;
            font-weight: 600;
            color: #059669;
            margin-bottom: 12px;
        }

        .due-date {
            font-size: 14px;
            color: #dc2626;
            font-weight: 600;
            margin-top: 12px;
        }

        .course-name {
            font-weight: 600;
            color: #6366f1;
        }

        .footer {
            font-size: 13px;
            color: #8b95b6;
            text-align: center;
            padding: 18px 24px 22px 24px;
            background: #f8fafc;
            border-bottom-left-radius: 18px;
            border-bottom-right-radius: 18px;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="header">
            📝 New Task Assigned
        </div>
        <div class="content">
            <p>You have a new task in <span class="course-name">%s</span>!</p>
            
            <div class="task-info">
                <div class="task-title">%s</div>
                <p>%s</p>
                <div class="due-date">📅 Due: %s</div>
            </div>
            
            <p>Head over to your dashboard to get started on this task. Don't forget to check the due date and requirements!</p>
            
            <p>Good luck with your studies!<br /><span style="color:#059669;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.CourseName, n.Title, n.Description, n.DueDate),
	}, nil
}
