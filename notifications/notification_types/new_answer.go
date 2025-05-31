package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// NewAnswerNotification represents a notification sent to teachers when a student submits an answer.
type NewAnswerNotification struct {
	TeacherName string `json:"teacher_name"`
	StudentName string `json:"student_name"`
	TaskTitle   string `json:"task_title"`
	CourseName  string `json:"course_name"`
	SubmittedAt string `json:"submitted_at"`
}

func (n *NewAnswerNotification) Type() string {
	return "NewAnswer"
}

func (n *NewAnswerNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *NewAnswerNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *NewAnswerNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "New Student Submission",
		Text:  n.StudentName + " submitted " + n.TaskTitle,
	}, nil
}

func (n *NewAnswerNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "New Submission: " + n.TaskTitle + " by " + n.StudentName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>New Student Submission</title>
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
            background: linear-gradient(90deg, #f59e0b 0%%, #fbbf24 100%%);
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

        .submission-info {
            background-color: #fffbeb;
            border-left: 4px solid #f59e0b;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .task-title {
            font-size: 20px;
            font-weight: 600;
            color: #d97706;
            margin-bottom: 12px;
        }

        .student-name {
            font-weight: 600;
            color: #6366f1;
        }

        .course-name {
            font-weight: 600;
            color: #059669;
        }

        .submitted-time {
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
            📋 New Submission
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>You have a new submission to review!</p>
            
            <div class="submission-info">
                <div class="task-title">%s</div>
                <p>Student: <span class="student-name">%s</span></p>
                <p>Course: <span class="course-name">%s</span></p>
                <div class="submitted-time">⏰ Submitted: %s</div>
            </div>
            
            <p>You can review the submission and provide feedback through your instructor dashboard.</p>
            
            <p>Happy teaching!<br /><span style="color:#f59e0b;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.TeacherName, n.TaskTitle, n.StudentName, n.CourseName, n.SubmittedAt),
	}, nil
}
