package advent2015

import (
	"testing"
	"my.app/pkg/advent2015"
)

func TestFirstPresentDimension(t *testing.T) {
	dimensions := "2x3x4"

	want := 58
	if got := advent2015.CalculatePresentsDimensions(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}

func TestSecondPresentDimension(t *testing.T) {
	dimensions := "1x1x10"

	want := 43
	if got := advent2015.CalculatePresentsDimensions(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}

func TestThirdPresentDimension(t *testing.T) {
	dimensions := "1x1x10\n1x1x10"

	want := 86
	if got := advent2015.CalculatePresentsDimensions(dimensions); got != want {
		t.Errorf("CalculatePresentsDimensions = %d, want %d", got, want)
	}
}