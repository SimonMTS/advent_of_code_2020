package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strconv"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	inputStrings := bp.GetN("day8/day8_input.txt")

	_, acc := runUntilRepeat(inputStrings)
	fmt.Println("\tpart 1:", acc)

	for j := 0; j < len(inputStrings); j++ {
		flipOP(j, &inputStrings)

		i, acc := runUntilRepeat(inputStrings)
		if i == len(inputStrings) {
			fmt.Println("\tpart 2:", acc)
			return
		}

		flipOP(j, &inputStrings)
	}

}

func runUntilRepeat(inputStrings []string) (i, accumulator int) {
	executed := make(map[int]string)
	i, accumulator = 0, 0
	for ; i < len(inputStrings); i++ {
		if _, ok := executed[i]; ok {
			break
		}

		executed[i] = inputStrings[i]
		parts := strings.Split(inputStrings[i], " ")
		if parts[0] == "acc" {
			inc, _ := strconv.Atoi(parts[1])
			accumulator += inc
		} else if parts[0] == "jmp" {
			inc, _ := strconv.Atoi(parts[1])
			i = (i - 1) + inc
		}
	}

	return
}

func flipOP(index int, inputStrings *[]string, recurse ...bool) {
	parts := strings.Split((*inputStrings)[index], " ")
	if parts[0] == "nop" {
		(*inputStrings)[index] = "jmp " + parts[1]
	} else if parts[0] == "jmp" {
		(*inputStrings)[index] = "nop " + parts[1]
	}
}
