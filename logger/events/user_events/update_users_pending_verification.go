package user_events

import "encoding/json"

type UpdateUsersPendingVerification struct {
	Number uint
}

func (e UpdateUsersPendingVerification) Type() string {
	return "UpdateUsersPendingVerification"
}

func (e UpdateUsersPendingVerification) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UpdateUsersPendingVerification")
}

func (e *UpdateUsersPendingVerification) Decode(data []byte) error {
	return Wrapp("error encoding UserBlocked", json.Unmarshal(data, e))
}
