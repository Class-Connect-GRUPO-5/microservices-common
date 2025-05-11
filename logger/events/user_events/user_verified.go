package user_events

import "encoding/json"

// After email verification
type UserVerified struct {
	UserID string `json:"user_id"`
}

func (e UserVerified) Type() string {
	return "UserVerified"
}

func (e UserVerified) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserVerified")
}

func (e *UserVerified) Decode(data []byte) error {
	return Wrapp("error encoding UserVerified", json.Unmarshal(data, e))
}
