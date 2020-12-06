package day5

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strconv"
)

const split_pos  = 7 // 13 / 7 / 18
const input_file = "day5/day5_input.txt" // ihs2 / day5_input / bigboy

func Run() {
	input_strings := bp.GetN(input_file)

	lowest := int(^uint(0) >> 1)
	highest := 0
	sum := 0
	fmt.Printf("\t%09d\n", len(input_strings))
	for i, s := range input_strings {
		id := codeToValues(s)
		sum += id
		if id < lowest { lowest = id }
		if id > highest { highest = id }

		if i % 100000 == 0 { fmt.Printf("\t%09d\r", i) }
	}

	expected_sum := 0
	for i := lowest; i <= highest; i++ { expected_sum += i }

	fmt.Print("\n\tpart1: "); fmt.Println(highest)
	fmt.Print("\tpart2: "); fmt.Println(expected_sum - sum)
}

func codeToValues(code string) int {
	res := ""
	for _, s := range code {
		if s == 'F' || s == 'L' { res += "0" } else { res += "1" }
	}

	id, _ := strconv.ParseInt(res, 2, 64)
	return int(id)
}
