package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/pangaunn/learn-go-1st/subject"
)

func init() {
	println("init naja")
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	// slice
	// slice := []string{}
	// slice = append(slice, "aa")
	// fmt.Println(slice)
	// slice := []string{"1", "2", "5"}
	// reverseSlice
	// fmt.Println(reverseSlice(slice))
	// fmt.Println(append(slice, 444))

	// map
	// mapVal := map[string]interface{}{}
	// mapVal["aaa"] = "3333"
	// mapVal["mapja"] = map[string]interface{}{}
	// fmt.Println(mapVal)

	// function
	// Println()
	// Println("Pang", "Kung", "na", "ja")
	// fmt.Println(Sum(5, 7))

	// sum, avg := sumAndAvg(2, 3, 4, 7, 9)
	// fmt.Println(sum, avg)
	// maStack := stack{1, 2, 3, 4, 5}
	// maStack = maStack.push(8)
	// fmt.Println(maStack)
	// maStack, _, _ = maStack.pop()
	// fmt.Println(maStack)

	fmt.Println(subject.GetSubject())
}

func Hello(name, surname string) {
	Println("Hello", name, surname)
}

func Println(a ...string) {
	for _, e := range a {
		println(e)
	}
}

func Sum(x, y int) int {
	return x + y
}

func sumAndAvg(number ...int) (int, float32) {
	sum := 0
	count := 0
	for _, n := range number {
		if isPrime(n) {
			sum += n
			count++
		}
	}
	avg := float32(sum) / float32(count)
	return sum, avg
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

type stack []int

func (s stack) push(a int) stack {
	s = append(s, a)
	return s
}

func (s stack) pop() (stack, int, error) {
	if len(s) == 0 {
		err := errors.New("stack dudeeeeeeeeeeee")
		return stack{}, 0, err
	}

	s = s[:len(s)-1]
	return s, s[len(s)-1], nil
}
