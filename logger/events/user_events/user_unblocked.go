package user_events

import "encoding/json"

type UserUnblocked struct {
	UserID string `json:"user_id"`
}

func (e UserUnblocked) Type() string {
	return "UserUnblocked"
}

func (e UserUnblocked) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserUnblocked")
}

func (e *UserUnblocked) Decode(data []byte) error {
	return Wrapp("error encoding UserUnblocked", json.Unmarshal(data, e))
}
