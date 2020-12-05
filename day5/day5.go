package main

import (
	"fmt"
    "math"
    "io/ioutil"
	"strings"
)

const split_pos = 18 // 13 / 7 / 18
const input_file = "./bigboy.txt" // ihs2 / day5_input / bigboy

func main() {
	bytes , _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(bytes),"\n")
	len_input_strings := len(input_strings)

	lowest := int(^uint(0) >> 1)
	highest := 0
	sum := 0
	fmt.Printf("%9d\n", len_input_strings)
	for i, s := range input_strings {
		_, _, id := codeToValues(s)
		sum += id
		if id < lowest { lowest = id }
		if id > highest { highest = id }

		if i % 100000 == 0 {
			fmt.Printf("%9d", i)
			fmt.Print("\r")
		}
	}

	expected_sum := 0
	for i := lowest; i <= highest; i++ {
		expected_sum += i
	}

	fmt.Print("\nhighest id: "); fmt.Println(highest)
	fmt.Print("my id: "); fmt.Println(expected_sum - sum)
}

func codeToValues(code string) (col, row, id int) {
	res := ""
	for _, s := range code {
		if s == 'F' || s == 'L' { res += "1" }
		if s == 'B' || s == 'R' { res += "0" }
	}

	row = binaryReduce(0, math.Pow(2,split_pos)-1, res[:split_pos])
	col = binaryReduce(0, math.Pow(2,float64(len(code)-split_pos))-1, res[split_pos:])
	id = (row * int(math.Pow(2,float64(len(code)-split_pos)))) + col
	return
}

func binaryReduce(lower, upper float64, path string) int {
	if len(path) == 0 { return int(upper) }

	if (path[0] == '1') {
		upper = upper - math.Ceil((upper - lower) / 2)
	} else {
		lower = lower + math.Ceil((upper - lower) / 2)
	}
	
	return int(binaryReduce(lower, upper, path[1:]))
}
