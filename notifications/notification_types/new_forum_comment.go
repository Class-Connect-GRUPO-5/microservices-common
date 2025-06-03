package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// NewForumCommentNotification represents a notification sent to teachers when a student submits an answer.
type NewForumCommentNotification struct {
	UserName       string `json:"user_name"`
	PostTitle      string `json:"post_title"`
	CommentContent string `json:"comment_content"`
}

func (n *NewForumCommentNotification) Type() string {
	return "NewForumComment"
}

func (n *NewForumCommentNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *NewForumCommentNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *NewForumCommentNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: n.UserName + " has commented on a post",
		Text:  "There is a new comment on the post: " + n.PostTitle,
	}, nil
}

func (n *NewForumCommentNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "New Comment on Post: " + n.PostTitle,
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>New Forum Comment</title>
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

        .comment-info {
            background-color: #eff6ff;
            border-left: 4px solid #3b82f6;
            padding: 20px;
            margin: 20px 0;
            border-radius: 8px;
        }

        .comment-content {
            background-color: #f8fafc;
            border: 1px solid #e2e8f0;
            border-radius: 8px;
            padding: 16px;
            margin: 16px 0;
            font-style: italic;
            color: #4a5568;
            line-height: 1.6;
        }

        .post-title {
            font-size: 20px;
            font-weight: 600;
            color: #1d4ed8;
            margin-bottom: 12px;
        }

        .user-name {
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
            💬 New Forum Comment
        </div>
        <div class="content">
            <p>There's a new comment on a forum post!</p>
            
            <div class="comment-info">
                <div class="post-title">%s</div>
                <p>Comment by: <span class="user-name">%s</span></p>
                <div class="comment-content">
                    "%s"
                </div>
            </div>
            
            <p>Check out the discussion and join the conversation in the forum.</p>
            
            <p>Happy learning!<br /><span style="color:#3b82f6;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.PostTitle, n.UserName, n.CommentContent),
	}, nil
}
