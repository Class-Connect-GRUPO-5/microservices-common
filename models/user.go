package models

import "time"

// User represents both the verified user and the pending user in the system.
// It contains the user's personal details, verification status, and role.
type User struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	VerifiedAt       *time.Time `json:"verified_at,omitempty"`
	Location         string     `json:"location,omitempty"`
	IsBlocked        bool       `json:"is_blocked"`
	RegistrationDate time.Time  `json:"registration_date"`
	Role             string     `json:"role"`
}

// UserToVerify represents a user who has registered but is still pending verification.
// This struct is used for pending users who need to verify their email before becoming active.
type UserToVerify struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Location     string    `json:"location"`
	Pin          string    `json:"pin,omitempty"`
	CreationTime time.Time `json:"created_at,omitempty"`
}

// UserProfile represents the user’s public profile data that can be updated by the user.
type UserProfile struct {
	UserID         string  `json:"user_id"`
	Name           string  `json:"name"`
	Location       *string `json:"location"`
	ProfilePicture *string `json:"profile_picture"`
	Biography      *string `json:"biography"`
}

// LoginAttempts tracks the login attempts for a user, including the number of failed attempts,
// the lockout time, and the number of times the user has been locked out.
type LoginAttempts struct {
	UserID         string     `json:"user_id"`
	FailedAttempts int        `json:"failed_attempts"`
	LockTime       *time.Time `json:"lock_time"`
	LockCount      int        `json:"lock_count"`
}
