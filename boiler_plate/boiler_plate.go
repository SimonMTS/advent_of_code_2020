package boilerplate

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Get gets input from file as single string
func Get(input string) string {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// GetC gets input from file and splits on ""
func GetC(input string) []string {
	return strings.Split(Get(input), "")
}

// GetN gets input from file and splits on "\n"
func GetN(input string) []string {
	return strings.Split(Get(input), "\n")
}

// GetNN gets input from file and splits on "\n\n"
func GetNN(input string) []string {
	return strings.Split(Get(input), "\n\n")
}

// Ntos changes all newlines to spaces
func Ntos(inputStrings []string) []string {
	for i := range inputStrings {
		for j := range inputStrings[i] {
			if string(inputStrings[i][j]) == "\n" {
				out := []rune(inputStrings[i])
				out[j] = ' '
				inputStrings[i] = string(out)
			}
		}
	}

	return inputStrings
}

// Atoi treats all lines as numbers, and converts them to ints
func Atoi(inputStrings []string) []int {
	inputInts := make([]int, len(inputStrings))
	for i, s := range inputStrings {
		asInt, _ := strconv.Atoi(s)
		inputInts[i] = asInt
	}

	return inputInts
}
