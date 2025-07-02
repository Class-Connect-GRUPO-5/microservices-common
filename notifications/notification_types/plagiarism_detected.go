package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// PlagiarismDetected represents a notification sent to teachers when a student's answer is detected for plagiarism.
type PlagiarismDetected struct {
	TeacherName       string  `json:"teacher_name"`
	StudentName       string  `json:"student_name"`
	TaskTitle         string  `json:"task_title"`
	CourseName        string  `json:"course_name"`
	SubmissionPreview string  `json:"submission_preview"`
	SimilarityScore   float64 `json:"similarity_score"`
	DetectedAt        string  `json:"detected_at"`
	MatchCount        int     `json:"match_count"`
}

func (n *PlagiarismDetected) Type() string {
	return "PlagiarismDetected"
}

func (n *PlagiarismDetected) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *PlagiarismDetected) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *PlagiarismDetected) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "🚨 Plagiarism Alert",
		Text:  fmt.Sprintf("%s flagged for plagiarism in %s (%.1f%% similarity)", n.StudentName, n.TaskTitle, n.SimilarityScore*100),
	}, nil
}

func (n *PlagiarismDetected) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "⚠️ Plagiarism Detected: " + n.TaskTitle + " by " + n.StudentName,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Plagiarism Detection Alert</title>
    <style>
        body {
            background-color: #f4f7fb;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 0;
        }

        .container {
            max-width: 580px;
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
            font-size: 16px;
            color: #22223b;
            line-height: 1.6;
        }

        .content p {
            margin: 16px 0;
        }

        .alert-box {
            background-color: #fef2f2;
            border: 2px solid #fecaca;
            border-radius: 12px;
            padding: 20px;
            margin: 24px 0;
            text-align: center;
        }

        .alert-text {
            color: #dc2626;
            font-weight: 600;
            font-size: 18px;
            margin-bottom: 8px;
        }

        .similarity-score {
            background: linear-gradient(90deg, #dc2626, #ef4444);
            color: white;
            padding: 8px 16px;
            border-radius: 20px;
            font-weight: 700;
            font-size: 16px;
            display: inline-block;
            margin-top: 8px;
        }

        .submission-info {
            background-color: #f8fafc;
            border: 1px solid #e2e8f0;
            border-radius: 12px;
            padding: 24px;
            margin: 24px 0;
        }

        .task-title {
            font-size: 20px;
            font-weight: 600;
            color: #1e293b;
            margin-bottom: 16px;
        }

        .info-row {
            display: flex;
            justify-content: space-between;
            margin: 12px 0;
            padding: 8px 0;
            border-bottom: 1px solid #e2e8f0;
        }

        .info-label {
            font-weight: 600;
            color: #64748b;
        }

        .info-value {
            font-weight: 500;
            color: #1e293b;
        }

        .student-name {
            color: #dc2626;
            font-weight: 600;
        }

        .course-name {
            color: #059669;
            font-weight: 600;
        }

        .preview-box {
            background-color: #fafafa;
            border-left: 4px solid #dc2626;
            padding: 16px;
            margin: 20px 0;
            border-radius: 0 8px 8px 0;
            font-family: 'Courier New', monospace;
            font-size: 14px;
            color: #374151;
            line-height: 1.5;
        }

        .preview-label {
            font-weight: 600;
            color: #dc2626;
            margin-bottom: 8px;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }

        .ai-powered {
            background: linear-gradient(135deg, #f0f9ff 0%%, #e0f2fe 100%%);
            border: 1px solid #0ea5e9;
            border-radius: 12px;
            padding: 16px;
            margin: 24px 0;
            text-align: center;
        }

        .ai-badge {
            background: linear-gradient(90deg, #0ea5e9, #0284c7);
            color: white;
            padding: 6px 12px;
            border-radius: 20px;
            font-weight: 600;
            font-size: 14px;
            display: inline-block;
            margin-bottom: 8px;
        }

        .ai-text {
            color: #0369a1;
            font-size: 14px;
            font-weight: 500;
        }

        .action-items {
            background-color: #f1f5f9;
            border-radius: 12px;
            padding: 20px;
            margin: 24px 0;
        }

        .action-title {
            font-weight: 600;
            color: #1e293b;
            margin-bottom: 12px;
            font-size: 16px;
        }

        .action-items ul {
            margin: 0;
            padding-left: 20px;
        }

        .action-items li {
            margin: 8px 0;
            color: #475569;
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
            🚨 Plagiarism Detection Alert
        </div>
        <div class="content">
            <p>Hi <strong>%s</strong>,</p>
            
            <div class="alert-box">
                <div class="alert-text">⚠️ High similarity detected in student submission</div>
                <div class="similarity-score">%.1f%%%% Similarity Score</div>
            </div>
            
            <p>Our AI-powered plagiarism detection system has flagged a submission that exceeds the similarity threshold and requires your immediate review:</p>
            
            <div class="submission-info">
                <div class="task-title">📋 %s</div>
                
                <div class="info-row">
                    <span class="info-label">Student:</span>
                    <span class="info-value student-name">%s</span>
                </div>
                
                <div class="info-row">
                    <span class="info-label">Course:</span>
                    <span class="info-value course-name">%s</span>
                </div>
                
                <div class="info-row">
                    <span class="info-label">Detected:</span>
                    <span class="info-value">%s</span>
                </div>
                
                <div class="info-row">
                    <span class="info-label">Matches Found:</span>
                    <span class="info-value">%d potential sources</span>
                </div>
                
            </div>
            
            <div class="preview-box">
                <div class="preview-label">📄 Submission Preview:</div>
                "%s"
            </div>
            
            <div class="ai-powered">
                <div class="ai-badge">🤖 Powered by ClassConnect AI</div>
                <div class="ai-text">Advanced machine learning algorithms for accurate plagiarism detection</div>
            </div>
            
            <div class="action-items">
                <div class="action-title">📝 Recommended Actions:</div>
                <ul>
                    <li>Review the detailed plagiarism report in your dashboard</li>
                    <li>Examine the highlighted matching content</li>
                    <li>Compare with the identified source materials</li>
                    <li>Contact the student to discuss the findings</li>
                    <li>Apply your institution's academic integrity policies</li>
                    <li>Document your decision and any actions taken</li>
                </ul>
            </div>
            
            <p>You can access the complete plagiarism analysis report through your instructor dashboard. The report includes detailed match comparisons, source identification, and confidence scores.</p>
            
            <p>Thank you for maintaining academic integrity in your courses.<br />
            <span style="color:#dc2626;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved. | AI-Powered Education Platform
        </div>
    </div>
</body>

</html>`, n.TeacherName, n.SimilarityScore*100, n.TaskTitle, n.StudentName, n.CourseName, n.DetectedAt, n.MatchCount, n.SubmissionPreview),
	}, nil
}
