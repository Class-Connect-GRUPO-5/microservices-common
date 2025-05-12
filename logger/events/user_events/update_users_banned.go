package user_events

import "encoding/json"

type UpdateUsersBanned struct {
	Number uint
}

func (e UpdateUsersBanned) Type() string {
	return "UpdateUsersBanned"
}

func (e UpdateUsersBanned) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UpdateUsersBanned")
}

func (e *UpdateUsersBanned) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
