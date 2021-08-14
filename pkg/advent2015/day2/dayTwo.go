package advent2015

import (
	"sort"
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
	intList := convertToIntList(numbers)
	l, w, h := intList[0], intList[1], intList[2]
	smallSide := getSmallSide(l, w, h)
	result := (2*l*w + 2*w*h + 2*h*l + smallSide)
	return result
}

func getSmallSide(l int, w int, h int) int {
	numbers := []int{l * w, w * h, h * l}
	sort.Ints(numbers)
	return numbers[0]
}

func CalculateSmallestDimension(dimensions string) int {
	newDimensions := strings.Split(dimensions, "\n")
	resultPresent := 0
	for _, present := range newDimensions {
		resultPresent += calculateSmallest(present)
	}
	return resultPresent
}

func calculateSmallest(dimensions string) int {
	numbers := strings.Split(dimensions, "x")
	intList := convertToIntList(numbers)
	l, w, h := intList[0], intList[1], intList[2]
	feetOfRibbon := calculateFeetOfRibbon(intList)
	ribbonForTheBow := (l * w * h)
	result := ribbonForTheBow + feetOfRibbon
	return result
}

func calculateFeetOfRibbon(numbers []int) int {
	sort.Ints(numbers)
	return sumSmallestSides(numbers[0], numbers[1])
}

func sumSmallestSides(side1 int, side2 int) int {
	return (side1*2 + side2*2)
}

func convertToIntList(list []string) []int {
	var intList = []int{}
	for _, number := range list {
		newNumber, _ := strconv.Atoi(number)
		intList = append(intList, newNumber)
	}
	return intList
}
