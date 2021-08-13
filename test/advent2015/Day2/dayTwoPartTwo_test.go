package advent2015

import (
	"io/ioutil"
	"testing"
	"my.app/pkg/advent2015"
)

func TestFirstWrapAPresent(t *testing.T) {
	dimensions := "2x3x4"

	want := 34
	if got := advent2015.CalculateSmallestDimension(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}

func TestSecondWrapAPresent(t *testing.T) {
	dimensions := "1x1x10"

	want := 14
	if got := advent2015.CalculateSmallestDimension(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}

func TestThirdWrapAPresent(t *testing.T) {
	dimensions := "1x1x10\n1x1x10"

	want := 28
	if got := advent2015.CalculateSmallestDimension(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}

func TestFinalWrapAPresent(t *testing.T) {
	input, _ := ioutil.ReadFile("./../../../resources/advent2015/dayTwo_input.txt")

	dimensions := string(input)

	want := 3737498
	if got := advent2015.CalculateSmallestDimension(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}