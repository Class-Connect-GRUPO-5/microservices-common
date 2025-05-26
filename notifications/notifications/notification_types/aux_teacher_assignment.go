package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
)

// AuxTeacherAssignmentNotification represents a notification sent when someone is assigned as an auxiliary teacher.
type AuxTeacherAssignmentNotification struct {
	TeacherName string `json:"teacher_name"`
	CourseName  string `json:"course_name"`
	MainTeacher string `json:"main_teacher"`
}

func (n *AuxTeacherAssignmentNotification) Type() string {
	return "AuxTeacherAssignment"
}

func (n *AuxTeacherAssignmentNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *AuxTeacherAssignmentNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *AuxTeacherAssignmentNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "You are now an Auxiliary Teacher!",
		Text:  "You've been assigned as aux teacher for " + n.CourseName,
	}, nil
}

func (n *AuxTeacherAssignmentNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "Auxiliary Teacher Assignment: " + n.CourseName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Auxiliary Teacher Assignment</title>
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
            background: linear-gradient(90deg, #dc2626 0%%, #ef4444 100%%);
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

        .assignment-info {
            background-color: #fef2f2;
            border-left: 4px solid #dc2626;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .course-name {
            font-size: 20px;
            font-weight: 600;
            color: #b91c1c;
            margin-bottom: 12px;
        }

        .main-teacher {
            font-weight: 600;
            color: #6366f1;
        }

        .assigned-time {
            font-size: 14px;
            color: #6b7280;
            margin-top: 12px;
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
            👨‍🏫 Teaching Assignment
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>Congratulations! You've been assigned as an auxiliary teacher for the following course:</p>
            
            <div class="assignment-info">
                <div class="course-name">%s</div>
                <p>Main Instructor: <span class="main-teacher">%s</span></p>
            </div>
            
            <p>You now have access to course materials, can assist with grading, and help support students. Check your instructor dashboard to get started!</p>
            
            <p>Thank you for your contribution to education!<br /><span style="color:#dc2626;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.TeacherName, n.CourseName, n.MainTeacher),
	}, nil
}
