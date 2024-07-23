package main

import (
	"os"
	"time"

	"github.com/cativovo/learn-go-with-tests/mocking"
)

func main() {
	sleeper := mocking.NewConfigurableSleeper(time.Sleep, time.Second*1)
	mocking.Countdown(os.Stdout, sleeper)
}
