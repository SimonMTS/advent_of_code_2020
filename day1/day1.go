package day1

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	bytes, _ := ioutil.ReadFile("day1/day1_input.txt") // real
	// bytes , _ := ioutil.ReadFile("./day1_input_big.txt")
	strings := strings.Split(string(bytes), "\n")

	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}

	sort.Ints(ints)

	var c1 chan string = make(chan string, 1)
	var c2 chan string = make(chan string, 1)
	part1(&ints, c1)
	part2(&ints, c2)

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}

func part1(ints *[]int, c chan string) {
	for i := 0; i < len(*ints); i++ {
		if check((*ints)[i], 0, ints) {
			c <- ("\tpart 1: " + strconv.Itoa((*ints)[i]*(2020-(*ints)[i])))
			close(c)
			return
		}
	}
}

func part2(ints *[]int, c chan string) {
	for i := 0; i < len(*ints); i++ {
		go func(i int, ints *[]int, c chan string) {
			for j := 0; j < len(*ints); j++ {
				if check((*ints)[i], (*ints)[j], ints) {
					c <- ("\tpart 2: " + strconv.Itoa((*ints)[i]*(*ints)[j]*(2020-((*ints)[i]+(*ints)[j]))))
				}
			}
		}(i, ints, c)
	}
}

func check(a int, b int, ints *[]int) bool {
	variable := a + b
	candidate := 2020 - variable
	found := sort.Search(len(*ints), func(j int) bool { return (*ints)[j] >= candidate })

	return found < len(*ints) && (*ints)[found] == candidate
}
