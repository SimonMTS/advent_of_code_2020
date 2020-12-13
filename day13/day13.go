package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strconv"
	"strings"
)

var dirs = []string{"north", "east", "south", "west"}

// Run is the pseudo main for this day
func Run() {
	input := bp.GetN("day13/day13_input.txt")
	time, _ := strconv.Atoi(input[0])
	busses := strings.Split(input[1], ",")
	indexedBusses := make([]int, len(busses)) // for part 2

	busID := 0
	lowest := time
	for i, bus := range busses {
		if bus == "x" {
			indexedBusses[i] = 0
			continue
		}
		departureInterval, _ := strconv.Atoi(bus)
		indexedBusses[i] = departureInterval

		waitTime := departureInterval - (time % departureInterval)
		if waitTime < lowest {
			lowest = waitTime
			busID = departureInterval
		}
	}

	fmt.Println("\tpart 1:", busID*lowest)

	// below was 'cheated' because I didn't have the math background
	// regarding the "Chinese remainder theorem"
	r, a := 0, 1
	for i, bus := range indexedBusses {
		if bus == 0 {
			continue
		}

		for (r+i)%bus != 0 {
			r += a
		}
		a *= bus
	}

	fmt.Println("\tpart 2:", r, "(cheated)")
}
