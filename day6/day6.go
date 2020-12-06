package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes , _ := ioutil.ReadFile("./day6_input.txt")
	input_strings := strings.Split(string(bytes),"\n\n")
	
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

	fmt.Print("\n sum part 1: "); fmt.Println(sum_p1)
	fmt.Print("\n sum part 2: "); fmt.Println(sum_p2)
}
