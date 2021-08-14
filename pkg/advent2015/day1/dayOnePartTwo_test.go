package advent2015

import "testing"

func TestSantasFloor_Part2(t *testing.T) {
	tests := []struct {
		name   string
		floors string
		want   int
	}{
		{
			name:   "Test 1: )",
			floors: ")",
			want:   1,
		}, {
			name:   "Test 2: ()())",
			floors: "()())",
			want:   5,
		}, {
			name:   "Test 3: (",
			floors: "(",
			want:   0,
		}, {
			name:   "Test 4: ReadFile",
			floors: GetInpuntFile(),
			want:   1783,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := FirstBasementWhereSantaHasBeen(test.floors); got != test.want {
				t.Errorf("FirstBasementWhereSantaHasBeen(floors) = %d, want %d", got, test.want)
			}
		})
	}
}
