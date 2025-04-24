package models

type Model interface {
	ToJSON() (string, error)
}
