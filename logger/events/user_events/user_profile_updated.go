package user_events

import "encoding/json"

type UserProfileUpdated struct {
	UserID string `json:"user_id"`
}

func (e UserProfileUpdated) Type() string {
	return "UserProfileUpdated"
}

func (e UserProfileUpdated) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserProfileUpdated")
}

func (e *UserProfileUpdated) Decode(data []byte) error {
	return Wrapp("error encoding UserProfileUpdated", json.Unmarshal(data, e))
}
