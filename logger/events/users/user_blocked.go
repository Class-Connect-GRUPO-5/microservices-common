package metrics

import "encoding/json"

type UserBlocked struct {
	UserID string `json:"user_id"`
}

func (e UserBlocked) Type() string {
	return "UserBlocked"
}

func (e UserBlocked) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserBlocked) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
