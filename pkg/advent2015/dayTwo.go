package advent2015

import (
	"strconv"
	"strings"
)

/*
CalculatePresentsDimensions is a function that calculate
total paper needed for christmas presents, part 2

https://adventofcode.com/2015/day/2
*/
func CalculatePresentsDimensions(dimensions string) int {
	newDimensions := strings.Split(dimensions, "\n")
	resultPresent := 0
	for _, present := range newDimensions {
		resultPresent += calculatePresentDimension(present)
	}
	return resultPresent
}

func calculatePresentDimension(dimensions string) int {
	numbers := strings.Split(dimensions, "x")
	l, w, h := convertNumbersToInt(numbers)
	smallSide := getSmallSide(l, w, h)
	result := (2*l*w + 2*w*h + 2*h*l + smallSide)
	return result
}

func convertNumbersToInt(numbers []string) (int, int, int) {
	l, _ := strconv.Atoi(numbers[0])
	w, _ := strconv.Atoi(numbers[1])
	h, _ := strconv.Atoi(numbers[2])
	return l, w, h
}

func getSmallSide(l int, w int, h int) int {
	side1 := l*w; side2 := w*h; side3 := h*l
	switch {
	case side1 < side2 && side1 < side3:
		return side1
	case side2 < side1 && side2 < side3:
		return side2
	default:
		return side3
	}
}