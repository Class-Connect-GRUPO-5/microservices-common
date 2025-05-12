package user_events

import "encoding/json"

// After email verification
type UserRecoveredPassword struct {
	UserID string `json:"user_id"`
}

func (e UserRecoveredPassword) Type() string {
	return "UserRecoveredPassword"
}

func (e UserRecoveredPassword) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserRecoveredPassword")
}

func (e *UserRecoveredPassword) Decode(data []byte) error {
	return Wrapp("error encoding UserRecoveredPassword", json.Unmarshal(data, e))
}
