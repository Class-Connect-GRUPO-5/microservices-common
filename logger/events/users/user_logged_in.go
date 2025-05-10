package metrics

import "encoding/json"

type UserLoggedIn struct {
	UserID string `json:"user_id"`
}

func (e UserLoggedIn) Type() string {
	return "UserLoggedIn"
}

func (e UserLoggedIn) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserLoggedIn) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
