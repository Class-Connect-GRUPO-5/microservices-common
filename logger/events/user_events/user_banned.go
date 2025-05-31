package user_events

import "encoding/json"

type UserBanned struct {
	UserID string `json:"user_id"`
}

func (e UserBanned) Type() string {
	return "UserBanned"
}

func (e UserBanned) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserBanned")
}

func (e *UserBanned) Decode(data []byte) error {
	return Wrapp("error encoding UserBanned", json.Unmarshal(data, e))
}
