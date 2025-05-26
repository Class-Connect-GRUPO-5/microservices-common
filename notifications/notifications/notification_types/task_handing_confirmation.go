package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notifications/notification_formats"
)

// TaskHandingConfirmationNotification represents a notification sent to students when they submit a task.
type TaskHandingConfirmationNotification struct {
	StudentName  string `json:"student_name"`
	TaskTitle    string `json:"task_title"`
	CourseName   string `json:"course_name"`
	SubmittedAt  string `json:"submitted_at"`
	SolutionText string `json:"solution_text"`
}

func (n *TaskHandingConfirmationNotification) Type() string {
	return "TaskHandingConfirmation"
}

func (n *TaskHandingConfirmationNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *TaskHandingConfirmationNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *TaskHandingConfirmationNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "Task Submitted Successfully",
		Text:  "Your submission for " + n.TaskTitle + " has been received",
	}, nil
}

func (n *TaskHandingConfirmationNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "Task Submission Confirmed: " + n.TaskTitle,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Task Submission Confirmed</title>
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
            background: linear-gradient(90deg, #10b981 0%%, #34d399 100%%);
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
            background-color: #ecfdf5;
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

        .course-name {
            font-weight: 600;
            color: #6366f1;
        }

        .submitted-time {
            font-size: 14px;
            color: #6b7280;
            margin-top: 12px;
        }

        .solution-text {
            background-color: #f9fafb;
            padding: 16px;
            border-radius: 8px;
            border: 1px solid #e5e7eb;
            margin-top: 12px;
            font-family: 'Courier New', monospace;
            color: #374151;
            white-space: pre-wrap;
            word-wrap: break-word;
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
            ✅ Submission Confirmed
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>Your task submission has been successfully received!</p>
            
            <div class="submission-info">
                <div class="task-title">%s</div>
                <p>Course: <span class="course-name">%s</span></p>
                <div class="submitted-time">⏰ Submitted: %s</div>
                <div class="solution-text">%s</div>
            </div>
            
            <p>Your instructor will review your submission and provide feedback. You can check for updates in your dashboard.</p>
            
            <p>Great work!<br /><span style="color:#10b981;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.StudentName, n.TaskTitle, n.CourseName, n.SubmittedAt, n.SolutionText),
	}, nil
}
