package user_events

import "encoding/json"

type UpdateUsersBlocked struct {
	Number uint
}

func (e UpdateUsersBlocked) Type() string {
	return "UpdateUsersBlocked"
}

func (e UpdateUsersBlocked) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UpdateUsersBlocked")
}

func (e *UpdateUsersBlocked) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
