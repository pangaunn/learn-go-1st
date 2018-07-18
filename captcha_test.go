package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type r interface {
	rand() int
}

type Range struct {
	start int
	end   int
}

type RangeMock struct {
	actualReturn int
}

func (r Range) rand() int {
	return random(r.start, r.end)
}

func (r RangeMock) rand() int {
	return r.actualReturn
}

// func main() {
// 	captchaString, result := generateCaptcha()
// 	fmt.Printf("%s = %d", captchaString, result)
// }

var timeJa int64 = 10

func random(min, max int) int {
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

func generateCaptcha(firstCaptchaR, secondCaptchaR, operatorCaptchaR r) (string, int) {
	firstCaptcha := firstCaptchaR.rand()
	secondCaptcha := secondCaptchaR.rand()

	result := 0
	operator := ""
	switch operationRandom := operatorCaptchaR.rand(); operationRandom {
	case 0:
		result = firstCaptcha + secondCaptcha
		operator = "+"
	case 1:
		result = firstCaptcha - secondCaptcha
		operator = "-"
	case 2:
		result = firstCaptcha * secondCaptcha
		operator = "*"
	default:
		result = firstCaptcha / secondCaptcha
		operator = "/"
	}

	captchaString := fmt.Sprintf("%d %s %d", firstCaptcha, operator, secondCaptcha)
	return captchaString, result
}

func TestCaptcha(t *testing.T) {
	t.Run("generateCaptcha", func(t *testing.T) {
		_, got := generateCaptcha(RangeMock{
			actualReturn: 15,
		}, RangeMock{
			actualReturn: 15,
		}, RangeMock{
			actualReturn: 0,
		})
		want := 30

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		_, got2 := generateCaptcha(RangeMock{
			actualReturn: 10,
		}, RangeMock{
			actualReturn: 5,
		}, RangeMock{
			actualReturn: 2,
		})
		want2 := 50

		if got2 != want2 {
			t.Errorf("got %d want %d", got2, want2)
		}
	})
}
