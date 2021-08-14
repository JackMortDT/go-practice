package advent2015

import (
	"io/ioutil"
	"testing"
)

func TestSantasFloor_Part1(t *testing.T) {
	tests := []struct {
		name   string
		floors string
		want   int
	}{
		{
			name:   "Test 1: (())",
			floors: "(())",
			want:   0,
		}, {
			name:   "Test 2: ()()",
			floors: "()()",
			want:   0,
		}, {
			name:   "Test 3: (((",
			floors: "(((",
			want:   3,
		}, {
			name:   "Test 4: (()(()(",
			floors: "(()(()(",
			want:   3,
		}, {
			name:   "Test 5: ))(((((",
			floors: "))(((((",
			want:   3,
		}, {
			name:   "Test 6: ())",
			floors: "())",
			want:   -1,
		}, {
			name:   "Test 7: ReadFile",
			floors: GetInpuntFile(),
			want:   232,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSantasFloor(tt.floors); got != tt.want {
				t.Errorf("CalculateSantasFloor(floors) = %d, want %d", got, tt.want)
			}
		})
	}
}

func GetInpuntFile() string {
	input, _ := ioutil.ReadFile("./../../../resources/advent2015/dayOne_input.txt")
	return string(input)
}
