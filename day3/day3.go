package main

import (
	"fmt"
	"strconv"
    "io/ioutil"
	"strings"
)

func main() {
	bytes , _ := ioutil.ReadFile("./day3_input.txt")
	input_strings := strings.Split(string(bytes),"")
	for i, c := range input_strings {
		if c == "\n" { input_strings = append(input_strings[:i], input_strings[i+1:]...) }
	}
	height := len(strings.Split(string(bytes),"\n"))
	width := len(strings.Split(string(bytes),"\n")[0])

	a1 := calcSlope(1, 1, height, width, &input_strings); fmt.Println("Right 1, down 1. :: " + strconv.Itoa(a1) + " trees")
	a2 := calcSlope(3, 1, height, width, &input_strings); fmt.Println("Right 3, down 1. :: " + strconv.Itoa(a2) + " trees // part 1")
	a3 := calcSlope(5, 1, height, width, &input_strings); fmt.Println("Right 5, down 1. :: " + strconv.Itoa(a3) + " trees")
	a4 := calcSlope(7, 1, height, width, &input_strings); fmt.Println("Right 7, down 1. :: " + strconv.Itoa(a4) + " trees")
	a5 := calcSlope(1, 2, height, width, &input_strings); fmt.Println("Right 1, down 2. :: " + strconv.Itoa(a5) + " trees")

	fmt.Println("All multiplied   :: " + strconv.Itoa(a1 * a2 * a3 * a4 * a5))
}

func calcSlope(Δx, Δy, height, width int, input_strings *[]string) int {
	trees := 0
	x := Δx
	y := Δy

	for y < height {
		if (*input_strings)[(x % width) + (y * width)] == "#" { trees++ }
		x += Δx
		y += Δy
	}

	return trees
}
