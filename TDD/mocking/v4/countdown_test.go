package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	sleepTimes = 4
	write      = "write"
	sleep      = "sleep"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep in-between the prints!", func(t *testing.T) {
		sco := &SpyCountdownOperations{}
		Countdown(sco, sco)
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, sco.Calls) {
			t.Errorf("wanted calls %v got %v", want, sco.Calls)
		}
	})
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
