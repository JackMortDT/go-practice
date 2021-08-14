package advent2015

import (
	"io/ioutil"
	"testing"
)

type Test struct {
	name       string
	dimensions string
	want       int
}

var test_day2_part1 = []Test{
	{"Test 1: 2x3x4", "2x3x4", 58},
	{"Test 2: 1x1x10", "1x1x10", 43},
	{"Test 3: 1x1x10\n1x1x10", "1x1x10\n1x1x10", 86},
	{"Test 4: ReadFile", getInpuntFile(), 1586300},
}

func TestPresentDimension_Part1(t *testing.T) {
	for _, test := range test_day2_part1 {
		t.Run(test.name, func(t *testing.T) {
			if got := CalculatePresentsDimensions(test.dimensions); got != test.want {
				t.Errorf("CalculatePresentsDimensions(dimensions) = %d, want %d", got, test.want)
			}
		})
	}
}

var test_day2_part2 = []Test{
	{"Test 1: 2x3x4", "2x3x4", 34},
	{"Test 2: 1x1x10", "1x1x10", 14},
	{"Test 3: 1x1x10\n1x1x10", "1x1x10\n1x1x10", 28},
	{"Test 4: ReadFile", getInpuntFile(), 3737498},
}

func TestPresentDimension_Part2(t *testing.T) {
	for _, test := range test_day2_part2 {
		t.Run(test.name, func(t *testing.T) {
			if got := CalculateSmallestDimension(test.dimensions); got != test.want {
				t.Errorf("CalculateSmallestDimension(dimensions) = %d, want %d", got, test.want)
			}
		})
	}
}

func getInpuntFile() string {
	input, _ := ioutil.ReadFile("./../../../resources/advent2015/dayTwo_input.txt")
	return string(input)
}
