package day1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	bytes, _ := ioutil.ReadFile("day2/day2_input.txt")
	inputStrings := strings.Split(string(bytes), "\n")

	part1(&inputStrings)
	part2(&inputStrings)
}

func part1(inputStrings *[]string) {
	validPasswords := 0

	for _, str := range *inputStrings {
		min, _ := strconv.Atoi(strings.Split(str, "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(str, "-")[1], " ")[0])
		letter := strings.Split(strings.Split(str, ":")[0], " ")[1]
		password := strings.TrimSpace(strings.Split(str, ":")[1])

		regex := regexp.MustCompile(letter)
		matches := regex.FindAllStringIndex(password, -1)

		if len(matches) <= max && len(matches) >= min {
			validPasswords++
		}
	}

	fmt.Println("\tpart 1:", strconv.Itoa(validPasswords))
}

func part2(inputStrings *[]string) {
	validPasswords := 0

	for _, str := range *inputStrings {
		min, _ := strconv.Atoi(strings.Split(str, "-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(str, "-")[1], " ")[0])
		letter := strings.Split(strings.Split(str, ":")[0], " ")[1]
		password := strings.TrimSpace(strings.Split(str, ":")[1])

		match1 := (len(password) > min-1 && string(password[min-1]) == letter)
		match2 := (len(password) > max-1 && string(password[max-1]) == letter)

		if match1 != match2 && (match1 || match2) {
			validPasswords++
		}
	}

	fmt.Println("\tpart 2:", strconv.Itoa(validPasswords))
}
