package day1

import (
	"fmt"
	"strconv"
    "io/ioutil"
	"strings"
	"regexp"
)

func Run() {
	bytes , _ := ioutil.ReadFile("day2/day2_input.txt")
	input_strings := strings.Split(string(bytes),"\n")

	part1(&input_strings);
	part2(&input_strings);
}

func part1(input_strings *[]string) {
	valid_passwords := 0

	for _, str := range *input_strings {
		min, _ := strconv.Atoi(strings.Split(str,"-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(str,"-")[1]," ")[0])
		letter := strings.Split(strings.Split(str,":")[0]," ")[1]
		password := strings.TrimSpace(strings.Split(str,":")[1])
		
		regex := regexp.MustCompile(letter)
		matches := regex.FindAllStringIndex(password, -1)

		if len(matches) <= max && len(matches) >= min {
			valid_passwords++
		}
	}

	fmt.Print("\tpart1: ")
	fmt.Println(strconv.Itoa(valid_passwords))
}

func part2(input_strings *[]string) {
	valid_passwords := 0

	for _, str := range *input_strings {
		min, _ := strconv.Atoi(strings.Split(str,"-")[0])
		max, _ := strconv.Atoi(strings.Split(strings.Split(str,"-")[1]," ")[0])
		letter := strings.Split(strings.Split(str,":")[0]," ")[1]
		password := strings.TrimSpace(strings.Split(str,":")[1])
		
		match1 := (len(password) > min-1 && string(password[min-1]) == letter)
		match2 := (len(password) > max-1 && string(password[max-1]) == letter)

		if match1 != match2 && (match1 || match2) {
			valid_passwords++
		}
	}

	fmt.Print("\tpart2: ")
	fmt.Println(strconv.Itoa(valid_passwords))
}
