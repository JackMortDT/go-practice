package advent2015

import (
	"io/ioutil"
	"testing"
	"my.app/pkg/advent2015"
)

func TestFirstBasementWhereSantaHasBeen(t *testing.T) {
	floors := ")"
	
	want := 1
	if got := advent2015.FirstBasementWhereSantaHasBeen(floors); got != want {
		t.Errorf("FirstBasementWhereSantaHasBeen() = %d, want %d", got, want)
	}
}

func TestSecondBasementWhereSantaHasBeen(t *testing.T) {
	floors := "()())"
	
	want := 5
	if got := advent2015.FirstBasementWhereSantaHasBeen(floors); got != want {
		t.Errorf("FirstBasementWhereSantaHasBeen() = %d, want %d", got, want)
	}
}

func TestFinalBasementWhereSantaHasBeen(t *testing.T) {
	input, _ := ioutil.ReadFile("./../../../resources/advent2015/dayOne_input.txt")

	floors := string(input)

	want := 1783
	if got := advent2015.FirstBasementWhereSantaHasBeen(floors); got != want {
		t.Errorf("FirstBasementWhereSantaHasBeen() = %d, want %d", got, want)
	}
}