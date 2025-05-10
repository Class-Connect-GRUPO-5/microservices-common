package user_events

import "encoding/json"

type AccountStatus int

const (
	AccountStatusRegistered AccountStatus = iota
	AcocuntStatusVerified
	AccountStatusNotExists
)

func (s AccountStatus) String() string {
	switch s {
	case AccountStatusRegistered:
		return "not verified"
	case AcocuntStatusVerified:
		return "verified"
	case AccountStatusNotExists:
		return "not registered"
	default:
		return "invalid"
	}
}

type UserFailedLogInAttempt struct {
	Email  string `json:"email"`
	Status string `json:"exists"`
}

func (e UserFailedLogInAttempt) Type() string {
	return "UserFailedLogInAttempt"
}

func (e UserFailedLogInAttempt) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserFailedLogInAttempt")
}

func (e *UserFailedLogInAttempt) Decode(data []byte) error {
	return Wrapp("error encoding UserFailedLogInAttempt", json.Unmarshal(data, e))
}
