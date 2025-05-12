package user_events

import "encoding/json"

// For first registration, before verifying their email
type AdminRegistered struct {
	UserID string `json:"user_id"`
}

func (e AdminRegistered) Type() string {
	return "AdminRegistered"
}

func (e AdminRegistered) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding AdminRegistered")
}

func (e *AdminRegistered) Decode(data []byte) error {
	return Wrapp("error encoding AdminRegistered", json.Unmarshal(data, e))
}
