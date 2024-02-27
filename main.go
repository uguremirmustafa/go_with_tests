package main

import (
	"os"
	"time"

	"github.com/uguremirmustafa/go_with_tests/mocks"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := mocks.NewConfigurableSleeper(400*time.Millisecond, time.Sleep)
	mocks.Countdown(os.Stdout, sleeper)
}
