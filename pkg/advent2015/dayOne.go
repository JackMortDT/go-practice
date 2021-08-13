package advent2015

/*
CalculateSantasFloor is a simple function
that calculate where is santa in the building part 1

https://adventofcode.com/2015/day/1
*/
func CalculateSantasFloor(floors string) int {

	floor := 0

	for _, char := range floors {
		if string(char) == "(" {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

/*
CalculateSantasFloor is a simple function
that calculate where is santa in the building part 2

https://adventofcode.com/2015/day/1
*/
func FirstBasementWhereSantaHasBeen(floors string) int {

	floor := 0
	basementNumber := 0

	for _, char := range floors {
		basementNumber++
		if string(char) == "(" {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return basementNumber
		}
	}

	return 1
}