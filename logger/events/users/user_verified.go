package metrics

import "encoding/json"

// After email verification
type UserVerified struct {
	UserID string `json:"user_id"`
}

func (e UserVerified) Type() string {
	return "UserVerified"
}

func (e UserVerified) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e *UserVerified) Decode(data []byte) error {
	return json.Unmarshal(data, e)
}
