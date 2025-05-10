package events

type Event interface {
	Type() string
	Encode() ([]byte, error)
	Decode([]byte) error
}
