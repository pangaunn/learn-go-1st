package fuzzbuzz

import "testing"

func Test1ShouldReturn1(t *testing.T) {
	expected := 1
	if SayFizBuzz(1) != expected {
		t.Error("noooooooooooooooo")
	}
}
func Test3ShouldReturnFizz(t *testing.T) {
	expected := "Fizz"
	if SayFizBuzz(3) != expected {
		t.Error("noooooooooooooooo")
	}
}

func Test5ShouldReturnBuzz(t *testing.T) {
	expected := "Buzz"
	if SayFizBuzz(5) != expected {
		t.Error("noooooooooooooooo")
	}
}

func Test15ShouldReturnFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	if SayFizBuzz(15) != expected {
		t.Error("noooooooooooooooo")
	}
}
