package user_events

import "encoding/json"

type UserLoggedIn struct {
	UserID string `json:"user_id"`
}

func (e UserLoggedIn) Type() string {
	return "UserLoggedIn"
}

func (e UserLoggedIn) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserLoggedIn")
}

func (e *UserLoggedIn) Decode(data []byte) error {
	return Wrapp("error encoding UserLoggedIn", json.Unmarshal(data, e))
}
