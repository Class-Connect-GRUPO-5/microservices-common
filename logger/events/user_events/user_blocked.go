package user_events

import "encoding/json"

type UserBlocked struct {
	UserID string `json:"user_id"`
}

func (e UserBlocked) Type() string {
	return "UserBlocked"
}

func (e UserBlocked) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserBlocked")
}

func (e *UserBlocked) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
