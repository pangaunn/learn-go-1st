package main

import (
	"fmt"
	"math"
)

func init() {
	println("init naja")
}

func main() {
	// slice
	// slice := []string{}
	// slice = append(slice, "aa")
	// fmt.Println(slice)

	// map
	// mapVal := map[string]interface{}{}
	// mapVal["aaa"] = "3333"
	// mapVal["mapja"] = map[string]interface{}{}
	// fmt.Println(mapVal)

	// function
	Println()
	Println("Pang", "Kung", "na", "ja")
	fmt.Println(Sum(5, 7))

	sum, avg := sumAndAvg(2, 3, 4, 7, 9)
	fmt.Println(sum, avg)
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
