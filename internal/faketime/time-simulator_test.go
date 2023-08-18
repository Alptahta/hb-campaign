package faketime

import (
	"fmt"
	"testing"
)

func Test_GetTime(t *testing.T) {
	expectedTime := "00:00"
	t.Run(fmt.Sprintf("Should return time as %s", expectedTime), func(t *testing.T) {
		ft := NewTime()
		if ft.GetTime() != expectedTime {
			t.Fatalf(fmt.Sprintf("expected %s but got %s", expectedTime, ft.GetTime()))
		}
	})
}

func Test_IncreaseTime(t *testing.T) {
	expectedFirstTime := "01:00"
	ft := NewTime()
	t.Run(fmt.Sprintf("Should return time as %s", expectedFirstTime), func(t *testing.T) {
		ft.IncreaseTime(1)
		if ft.GetTime() != expectedFirstTime {
			t.Fatalf(fmt.Sprintf("expected %s but got %s", expectedFirstTime, ft.GetTime()))
		}
	})
	expectedSecondTime := "03:00"
	t.Run(fmt.Sprintf("Should return time as %s", expectedSecondTime), func(t *testing.T) {
		ft.IncreaseTime(2)
		if ft.GetTime() != expectedSecondTime {
			t.Fatalf(fmt.Sprintf("expected %s but got %s", expectedSecondTime, ft.GetTime()))
		}
	})
}
