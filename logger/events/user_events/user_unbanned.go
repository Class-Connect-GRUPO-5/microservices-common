package user_events

import "encoding/json"

type UserUnbanned struct {
	UserID string `json:"user_id"`
}

func (e UserUnbanned) Type() string {
	return "UserUnbanned"
}

func (e UserUnbanned) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserUnbanned")
}

func (e *UserUnbanned) Decode(data []byte) error {
	return Wrapp("error encoding UserUnbanned", json.Unmarshal(data, e))
}
