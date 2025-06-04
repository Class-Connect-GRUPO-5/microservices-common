package notification_types

import (
	"encoding/json"
	"fmt"

	"github.com/Class-Connect-GRUPO-5/microservices-common/notifications/notification_formats"
)

// RulesUpdateNotification represents a notification sent to users when the application's terms and conditions have been updated.
type RulesUpdateNotification struct {
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

func (n *RulesUpdateNotification) Type() string {
	return "RulesUpdate"
}

func (n *RulesUpdateNotification) Encode() ([]byte, error) {
	return json.Marshal(n)
}

func (n *RulesUpdateNotification) Decode(data []byte) error {
	return json.Unmarshal(data, n)
}

func (n *RulesUpdateNotification) AsPush() (notification_formats.PushNotification, error) {
	return notification_formats.PushNotification{
		Title: "Terms and Conditions Updated",
		Text:  fmt.Sprintf("Hi %s, our terms and conditions have been updated (%s).", n.Name, n.UpdatedAt),
	}, nil
}

func (n *RulesUpdateNotification) AsEmail() (notification_formats.Email, error) {
	return notification_formats.Email{
		Subject: "We've Updated Our Terms and Conditions",
		Body: fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Terms and Conditions Update</title>
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

        .btn {
            display: inline-block;
            margin-top: 24px;
            padding: 12px 24px;
            background-color: #6366f1;
            color: white;
            text-decoration: none;
            border-radius: 6px;
            font-weight: 500;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="header">
            📋 Terms and Conditions Update
        </div>
        <div class="content">
            <p>Hi %s,</p>
            
            <p>We want to inform you that our <b>Terms and Conditions</b> have been updated on: <strong>%s</strong>.</p>
            
            <p>It's important that you review the changes to stay informed about your rights and responsibilities. You can check the new terms in the ClassConnect app.</p>
            
            <p>If you have any questions or need more information, don't hesitate to reply to this email or visit our Help Center.</p>
            
            <p>Thank you for trusting us.<br /><span style="color:#6366f1;font-weight:500;">The ClassConnect Team</span></p>
        </div>
        <div class="footer">
            &copy; 2025 ClassConnect. All rights reserved.
        </div>
    </div>
</body>

</html>`, n.Name, n.UpdatedAt),
	}, nil
}
