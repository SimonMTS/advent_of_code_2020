package day3

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	bytes, _ := ioutil.ReadFile("day3/day3_input.txt")
	inputStrings := strings.Split(string(bytes), "")
	for i, c := range inputStrings {
		if c == "\n" {
			inputStrings = append(inputStrings[:i], inputStrings[i+1:]...)
		}
	}
	height := len(strings.Split(string(bytes), "\n"))
	width := len(strings.Split(string(bytes), "\n")[0])

	a1 := calcSlope(1, 1, height, width, &inputStrings)
	fmt.Println("\tRight 1, down 1. :: " + strconv.Itoa(a1) + " trees")
	a2 := calcSlope(3, 1, height, width, &inputStrings)
	fmt.Println("\tRight 3, down 1. :: " + strconv.Itoa(a2) + " trees // part 1")
	a3 := calcSlope(5, 1, height, width, &inputStrings)
	fmt.Println("\tRight 5, down 1. :: " + strconv.Itoa(a3) + " trees")
	a4 := calcSlope(7, 1, height, width, &inputStrings)
	fmt.Println("\tRight 7, down 1. :: " + strconv.Itoa(a4) + " trees")
	a5 := calcSlope(1, 2, height, width, &inputStrings)
	fmt.Println("\tRight 1, down 2. :: " + strconv.Itoa(a5) + " trees")

	fmt.Println("\tAll multiplied   :: " + strconv.Itoa(a1*a2*a3*a4*a5) + " // part 2")
}

func calcSlope(Δx, Δy, height, width int, inputStrings *[]string) int {
	trees := 0
	x := Δx
	y := Δy

	for y < height {
		if (*inputStrings)[(x%width)+(y*width)] == "#" {
			trees++
		}
		x += Δx
		y += Δy
	}

	return trees
}
