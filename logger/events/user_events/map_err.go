package user_events

import "fmt"

type errMapper[T any] struct {
	val T
	err error
}

func Map[T any](val T, err error) errMapper[T] {
	return errMapper[T]{val, err}
}

func (m errMapper[T]) Err(prefix string) (T, error) {
	if m.err != nil {
		return m.val, fmt.Errorf("%s: %v", prefix, m.err)
	}
	return m.val, nil
}

func Wrapp(prefix string, err error) error {
	if err != nil {
		fmt.Errorf("%s: %v", prefix, err)
	}
	return nil
}
