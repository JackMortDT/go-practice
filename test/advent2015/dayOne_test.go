package advent2015

import (
	"io/ioutil"
	"testing"
	"my.app/pkg/advent2015"
)

func TestFirstSantasFloor(t *testing.T) {
	floors := "(())"

	want := 0
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}
}

func TestSecondSantasFloor(t *testing.T) {
	floors := "()()"

	want := 0
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}
}

func TestThirdSantasFloor(t *testing.T) {
	floors := "((("

	want := 3
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}
}

func TestFourthSantasFloor(t *testing.T) {
	floors := "(()(()("

	want := 3
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}
}

func TestFifthSantasFloor(t *testing.T) {
	floors := "))((((("

	want := 3
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}	
}

func TestSixthSantasFloor(t *testing.T) {
	floors := "())"

	want := -1
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}	
}

func TestFinalSantasFloor(t *testing.T) {
	input, err := ioutil.ReadFile("./../../resources/advent2015/dayOne_input.txt")
	if err != nil {
		panic(err)
	}

	floors := string(input)

	want := 232
	if got := advent2015.CalculateSantasFloor(floors); got != want {
		t.Errorf("CalculateSantasFloor = %d, want %d", got, want)
	}	
}