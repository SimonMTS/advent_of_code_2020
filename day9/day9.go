package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
)

// Run is the pseudo main for this day
func Run() {
	inputInts := bp.Atoi(bp.GetN("day9/day9_input.txt"))
	msg := inputInts[25:]

	i, number := 0, 0
	for i, number = range msg {
		preamble := inputInts[i : i+25]

		if !containsSum(preamble, number) {
			fmt.Println("\tpart 1:", number)
			break
		}
	}

	for index := range inputInts {
		if res := contiguousExistsFrom(index, number, inputInts); res != -1 {
			fmt.Println("\tpart 2:", res)
			break
		}
	}
}

func containsSum(preamble []int, number int) bool {
	for _, preambleNumber := range preamble {
		otherPreambleNumber := number - preambleNumber
		for _, subPreambleNumber := range preamble {
			if subPreambleNumber == otherPreambleNumber {
				return true
			}
		}
	}
	return false
}

func contiguousExistsFrom(startingIndex, sum int, list []int) int {
	runningCount := 0
	for i := startingIndex; i < len(list); i++ {
		runningCount += list[i]

		if runningCount == sum {
			smallest := list[startingIndex]
			for o := startingIndex; o <= i; o++ {
				if smallest > list[o] {
					smallest = list[o]
				}
			}

			largest := list[startingIndex]
			for o := startingIndex; o <= i; o++ {
				if largest < list[o] {
					largest = list[o]
				}
			}

			return smallest + largest
		}
	}

	return -1
}
