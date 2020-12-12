package day6

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	inputStrings := bp.GetNN("day6/day6_input.txt")

	sumP1, sumP2 := 0, 0
	for _, group := range inputStrings {
		qmap := make(map[string]int)
		peopleStrings := strings.Split(group, "\n")

		for _, person := range peopleStrings {
			questionStrings := strings.Split(person, "")

			for _, question := range questionStrings {
				if _, ok := qmap[question]; !ok {
					qmap[question] = 1
				} else {
					qmap[question]++
				}
			}
		}

		for _, val := range qmap {
			if val == len(peopleStrings) {
				sumP2++
			}
		}
		sumP1 += len(qmap)
	}

	fmt.Println("\tpart 1:", sumP1)
	fmt.Println("\tpart 2:", sumP2)
}
