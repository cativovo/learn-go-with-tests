package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := new(bytes.Buffer)
		spyCountdownOperations := new(SpyCountdownOperations)

		Countdown(buf, spyCountdownOperations)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("print then sleep", func(t *testing.T) {
		spyCountdownOperations := new(SpyCountdownOperations)

		Countdown(spyCountdownOperations, spyCountdownOperations)

		got := spyCountdownOperations.Calls
		want := []string{
			write, // 3
			sleep,
			write, // 2
			sleep,
			write, // 1
			sleep,
			write, // Go!
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5

	spyTime := new(SpyTime)
	sleeper := NewConfigurableSleeper(spyTime.Sleep, sleepTime)
	sleeper.Sleep()

	got := spyTime.durationSlept
	want := sleepTime

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
