package day6

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strings"
)

func Run() {
	input_strings := bp.GetNN("day6/day6_input.txt")
	
	sum_p1 := 0
	sum_p2 := 0
	for _, group := range input_strings {
		qmap := make(map[string]int)
		people_strings := strings.Split(group,"\n")

		for _, person := range people_strings {
			question_strings := strings.Split(person,"")
			
			for _, question := range question_strings {
				if _, ok := qmap[question]; !ok {
					qmap[question] = 1
				} else {
					qmap[question]++
				}
			}
		}

		for _, val := range qmap {
			if (val == len(people_strings)) { sum_p2++ }
		}
		sum_p1 += len(qmap)
	}

	fmt.Print("\tsum part 1: "); fmt.Println(sum_p1)
	fmt.Print("\tsum part 2: "); fmt.Println(sum_p2)
}
