package metrics

import "encoding/json"

type UserUnblocked struct {
	UserID string `json:"user_id"`
}

func (e UserUnblocked) Type() string {
	return "UserUnblocked"
}

func (e UserUnblocked) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserUnblocked) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
