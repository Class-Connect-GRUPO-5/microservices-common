package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
)

// TaskFeedbackNotification represents a notification sent to students when they receive feedback on their task.
type TaskFeedbackNotification struct {
	StudentName string `json:"student_name"`
	TaskTitle   string `json:"task_title"`
	CourseName  string `json:"course_name"`
	TeacherName string `json:"teacher_name"`
	Grade       string `json:"grade"`
	Feedback    string `json:"feedback"`
}

func (n *TaskFeedbackNotification) Type() string {
	return "TaskFeedback"
}

func (n *TaskFeedbackNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *TaskFeedbackNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *TaskFeedbackNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: n.TeacherName + " provided feedback on your task",
		Text:  "You received feedback for " + n.TaskTitle,
	}, nil
}

func (n *TaskFeedbackNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "Task Feedback: " + n.TaskTitle + " - " + n.CourseName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Task Feedback Received</title>
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
            background: linear-gradient(90deg, #8b5cf6 0%%, #a78bfa 100%%);
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

        .feedback-info {
            background-color: #faf5ff;
            border-left: 4px solid #8b5cf6;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .task-title {
            font-size: 20px;
            font-weight: 600;
            color: #7c3aed;
            margin-bottom: 12px;
        }

        .grade {
            font-size: 18px;
            font-weight: 600;
            color: #059669;
            margin: 12px 0;
        }

        .feedback-text {
            background-color: #f9fafb;
            padding: 16px;
            border-radius: 8px;
            border: 1px solid #e5e7eb;
            margin-top: 12px;
            font-style: italic;
            color: #374151;
        }

        .course-name {
            font-weight: 600;
            color: #6366f1;
        }

        .teacher-name {
            font-weight: 600;
            color: #059669;
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
            📝 Task Feedback
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>You've received feedback on your task submission!</p>
            
            <div class="feedback-info">
                <div class="task-title">%s</div>
                <p>Course: <span class="course-name">%s</span></p>
                <p>Instructor: <span class="teacher-name">%s</span></p>
                <div class="grade">📊 Grade: %s</div>
                <div class="feedback-text">%s</div>
            </div>
            
            <p>You can view the detailed feedback and your graded submission in your dashboard.</p>
            
            <p>Keep up the great work!<br /><span style="color:#8b5cf6;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.StudentName, n.TaskTitle, n.CourseName, n.TeacherName, n.Grade, n.Feedback),
	}, nil
}
