package day5

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strconv"
)

const splitPos = 7                      // 13 / 7 / 18
const inputFile = "day5/day5_input.txt" // ihs2 / day5_input / bigboy

// Run is the pseudo main for this day
func Run() {
	inputStrings := bp.GetN(inputFile)

	lowest := int(^uint(0) >> 1)
	highest := 0
	sum := 0
	fmt.Printf("\t%09d\n", len(inputStrings))
	for i, s := range inputStrings {
		id := codeToValues(s)
		sum += id
		if id < lowest {
			lowest = id
		}
		if id > highest {
			highest = id
		}

		if i%100000 == 0 {
			fmt.Printf("\t%09d\r", i)
		}
	}

	expectedSum := 0
	for i := lowest; i <= highest; i++ {
		expectedSum += i
	}

	fmt.Println("\n\tpart1:", highest)
	fmt.Println("\tpart2:", expectedSum-sum)
}

func codeToValues(code string) int {
	res := ""
	for _, s := range code {
		if s == 'F' || s == 'L' {
			res += "0"
		} else {
			res += "1"
		}
	}

	id, _ := strconv.ParseInt(res, 2, 64)
	return int(id)
}
