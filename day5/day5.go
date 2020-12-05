package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const split_pos  = 18 				// 13 	/ 7 		 / 18
const input_file = "./bigboy.txt" 	// ihs2 / day5_input / bigboy

func main() {
	bytes , _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(bytes),"\n")

	lowest := int(^uint(0) >> 1)
	highest := 0
	sum := 0
	fmt.Printf("%9d\n", len(input_strings))
	for i, s := range input_strings {
		id := codeToValues(s)
		sum += id
		if id < lowest { lowest = id }
		if id > highest { highest = id }

		if i % 100000 == 0 { fmt.Printf("%9d\r", i) }
	}

	expected_sum := 0
	for i := lowest; i <= highest; i++ { expected_sum += i }

	fmt.Print("\n highest id: "); fmt.Println(highest)
	fmt.Print(" missing id: "); fmt.Println(expected_sum - sum)
}

func codeToValues(code string) int {
	res := ""
	for _, s := range code {
		if s == 'F' || s == 'L' { res += "0" } else { res += "1" }
	}

	id, _ := strconv.ParseInt(res, 2, 64)
	return int(id)
}
