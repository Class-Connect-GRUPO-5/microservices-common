package user_events

import "encoding/json"

type UpdateUsersActive struct {
	Number uint
}

func (e UpdateUsersActive) Type() string {
	return "UpdateUsersActive"
}

func (e UpdateUsersActive) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UpdateUsersActive")
}

func (e *UpdateUsersActive) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
