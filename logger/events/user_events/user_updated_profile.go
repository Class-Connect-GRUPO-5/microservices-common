package user_events

import "encoding/json"

type UpdateUserProfile struct {
	UserID string `json:"user_id"`
}

func (e UpdateUserProfile) Type() string {
	return "UpdateUserProfile"
}

func (e UpdateUserProfile) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UpdateUserProfile")
}

func (e *UpdateUserProfile) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
