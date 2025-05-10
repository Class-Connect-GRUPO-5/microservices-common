package user_events

import "encoding/json"

type UserFailedLogInAttempt struct {
	Email  string `json:"email"`
	Exists bool   `json:"exists"`
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
