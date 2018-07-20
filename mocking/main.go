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

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (s *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, &RealSleeper{})
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
