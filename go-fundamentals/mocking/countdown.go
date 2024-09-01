package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

type ConfigurableSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func NewConfigurableSleeper(sleep func(time.Duration), duration time.Duration) ConfigurableSleeper {
	return ConfigurableSleeper{
		sleep:    sleep,
		duration: duration,
	}
}

const (
	FinalWord      = "Go!"
	CountdownStart = 3
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}
