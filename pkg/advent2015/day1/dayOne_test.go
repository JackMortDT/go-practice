package advent2015

import (
	"io/ioutil"
	"testing"
)

type Test struct {
	name   string
	floors string
	want   int
}

var tests_day1_part1 = []Test{
	{"Test 1: (())", "(())", 0},
	{"Test 2: ()()", "()()", 0},
	{"Test 3: (((", "(((", 3},
	{"Test 4: (()(()(", "(()(()(", 3},
	{"Test 5: ))(((((", "))(((((", 3},
	{"Test 6: ())", "())", -1},
	{"Test 7: ReadFile", getInpuntFile(), 232},
}

func TestSantasFloor_Part1(t *testing.T) {
	for _, test := range tests_day1_part1 {
		t.Run(test.name, func(t *testing.T) {
			if got := CalculateSantasFloor(test.floors); got != test.want {
				t.Errorf("CalculateSantasFloor(floors) = %d, want %d", got, test.want)
			}
		})
	}
}

var tests_day1_part2 = []Test{
	{"Test 1: )", ")", 1},
	{"Test 2: ()())", "()())", 5},
	{"Test 3: (", "(", 0},
	{"Test 4: ReadFile", getInpuntFile(), 1783},
}

func TestSantasFloor_Part2(t *testing.T) {
	for _, test := range tests_day1_part2 {
		t.Run(test.name, func(t *testing.T) {
			if got := FirstBasementWhereSantaHasBeen(test.floors); got != test.want {
				t.Errorf("FirstBasementWhereSantaHasBeen(floors) = %d, want %d", got, test.want)
			}
		})
	}
}

func getInpuntFile() string {
	input, _ := ioutil.ReadFile("./../../../resources/advent2015/dayOne_input.txt")
	return string(input)
}
