package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// InscriptionConfirmationNotification represents a notification sent to users when they successfully enroll in a course.
type InscriptionConfirmationNotification struct {
	StudentName string `json:"student_name"`
	CourseName  string `json:"course_name"`
}

func (n *InscriptionConfirmationNotification) Type() string {
	return "InscriptionConfirmation"
}

func (n *InscriptionConfirmationNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *InscriptionConfirmationNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *InscriptionConfirmationNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "Course Enrollment Confirmed",
		Text:  "You've successfully enrolled in " + n.CourseName,
	}, nil
}

func (n *InscriptionConfirmationNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "Course Enrollment Confirmed: " + n.CourseName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Course Enrollment Confirmed</title>
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
            background: linear-gradient(90deg, #3b82f6 0%%, #60a5fa 100%%);
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

        .course-info {
            background-color: #eff6ff;
            border-left: 4px solid #3b82f6;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .course-name {
            font-size: 20px;
            font-weight: 600;
            color: #1e40af;
            margin-bottom: 12px;
        }

        .teacher-name {
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
            🎓 Enrollment Confirmed
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>Great news! You've successfully enrolled in:</p>
            
            <div class="course-info">
                <div class="course-name">%s</div>
            </div>
            
            <p>You can now access course materials, assignments, and join class discussions. Check your dashboard to get started!</p>
            
            <p>Best of luck with your studies!<br /><span style="color:#3b82f6;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.StudentName, n.CourseName),
	}, nil
}
