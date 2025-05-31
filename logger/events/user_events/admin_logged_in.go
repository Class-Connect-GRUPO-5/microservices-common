package user_events

import "encoding/json"

type AdminLoggedIn struct {
	UserID string `json:"user_id"`
}

func (e AdminLoggedIn) Type() string {
	return "AdminLoggedIn"
}

func (e AdminLoggedIn) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding AdminLoggedIn")
}

func (e *AdminLoggedIn) Decode(data []byte) error {
	return Wrapp("error encoding AdminLoggedIn", json.Unmarshal(data, e))
}
