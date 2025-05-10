package user_events

import "encoding/json"

// For first registration, before verifying their email
type UserRegistered struct {
	UserID string `json:"user_id"`
}

func (e UserRegistered) Type() string {
	return "UserRegistered"
}

func (e UserRegistered) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserRegistered) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
