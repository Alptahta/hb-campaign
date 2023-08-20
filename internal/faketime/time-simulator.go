package faketime

import "fmt"

type TimeInterface interface {
	GetTime() string
	IncreaseTime(int)
}

type Time struct {
	Hour   int
	Minute int
}

func NewTime() Time {
	return Time{}
}

func (t *Time) GetTime() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}

func (t *Time) IncreaseTime(h int) {
	t.Hour = t.Hour + h
}
