package faketime

import "fmt"

type TimeInterface interface {
	GetTime() string
	IncreaseTime(int)
	NotifyAll()
	Register(o Observer)
}

type Time struct {
	Hour         int
	Minute       int
	observerList []Observer
}

func NewTime() Time {
	return Time{}
}

func (t *Time) GetTime() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}

func (t *Time) IncreaseTime(h int) {
	t.Hour = t.Hour + h
	t.NotifyAll()
}

func (t *Time) Register(o Observer) {
	t.observerList = append(t.observerList, o)
}

func (t *Time) NotifyAll() {
	for _, observer := range t.observerList {
		observer.Update()
	}
}
