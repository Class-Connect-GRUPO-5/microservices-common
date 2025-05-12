package user_events

import "encoding/json"

// After email verification
type UserStartedPasswordRecovery struct {
	UserEmail string `json:"email"`
}

func (e UserStartedPasswordRecovery) Type() string {
	return "UserStartedPasswordRecovery"
}

func (e UserStartedPasswordRecovery) Encode() ([]byte, error) {
	return Map(json.Marshal(e)).Err("error encoding UserStartedPasswordRecovery")
}

func (e *UserStartedPasswordRecovery) Decode(data []byte) error {
	return Wrapp("error encoding UserStartedPasswordRecovery", json.Unmarshal(data, e))
}
