package advent2015

/*
CalculateSantasFloor is a simple function
that calculate where is santa in the building

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
