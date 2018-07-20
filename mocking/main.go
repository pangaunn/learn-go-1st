package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

type Sleeper interface {
	Sleep()
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const coundownStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := coundownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}

	s.Sleep()
	fmt.Fprintf(w, finalWord)
}
