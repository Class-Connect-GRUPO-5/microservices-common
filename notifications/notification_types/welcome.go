package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// WelcomeNotification represents a notification sent to users when they first join the platform.
type WelcomeNotification struct {
	Name string `json:"name"`
}

func (n *WelcomeNotification) Type() string {
	return "Welcome"
}

func (n *WelcomeNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *WelcomeNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *WelcomeNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "Welcome to Class Connect",
		Text:  "Hello " + n.Name + ", welcome to Class Connect!",
	}, nil
}

func (n *WelcomeNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "Welcome to ClassConnect!",
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Welcome to ClassConnect</title>
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
            background: linear-gradient(90deg, #6366f1 0%%, #818cf8 100%%);
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
            Welcome to ClassConnect!
        </div>
        <div class="content">
            <p>Hi %s,</p>
            <p>We're thrilled to welcome you to <b>ClassConnect</b>! Your account is ready, and your learning journey
                begins now.</p>
            <p>Explore your dashboard, connect with classmates, and join your first class. If you need help, just reply
                to this email or visit our Help Center.</p>
            <p>Happy learning!<br /><span style="color:#6366f1;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.Name),
	}, nil
}
