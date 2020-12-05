package main

import (
	"fmt"
	"strconv"
    "io/ioutil"
	"strings"
	"regexp"
	"regexp2"
)

func main() {
	bytes , _ := ioutil.ReadFile("./day4_input.txt")
	input_strings := strings.Split(string(bytes),"\n\n")
	for i, _ := range input_strings { for j, _ := range input_strings[i] {
		if string(input_strings[i][j]) == "\n" {
			out := []rune(input_strings[i])
			out[j] = ' '
			input_strings[i] = string(out);
		}
	} }

	valid := 0
	for _, s := range input_strings {
		if isValid(s) { valid++ }
	}
	fmt.Println(valid)
}

func isValid(s string) bool {
	byr, err := strconv.Atoi(regexMatch(s, "(?<=byr:)\\d{4}"))
	if err != nil || byr < 1920 || byr > 2002 { return false }

	iyr, err := strconv.Atoi(regexMatch(s, "(?<=iyr:)\\d{4}"))
	if err != nil || iyr < 2010 || iyr > 2020 { return false }

	eyr, err := strconv.Atoi(regexMatch(s, "(?<=eyr:)\\d{4}"))
	if err != nil || eyr < 2020 || eyr > 2030 { return false }

	hgt_suffix := regexMatch(s, "(?<=hgt:\\d{2,3})\\D\\D")
	hgt, err := strconv.Atoi(regexMatch(s, "(?<=hgt:)\\d{2,3}"))
	if err != nil || 
		(hgt_suffix == "cm" && (hgt < 150 || hgt > 193)) || 
		(hgt_suffix == "in" && (hgt < 59 || hgt > 76)) ||
		(hgt_suffix != "cm" && hgt_suffix != "in") { return false }

	hcl := regexMatch(s, "hcl:#[0-9a-fA-F]{6}")
	if len(hcl) != 11 { return false }

	ecl := regexMatch(s, "ecl:(amb|blu|brn|gry|grn|hzl|oth)")
	if len(ecl) != 7 { return false }

	pid, err := strconv.Atoi(regexMatch(s, "(?<=pid:)\\d{9}(?=\\D|$)"))
	if err != nil || pid < 0 { return false }

	return true
}

func regexMatch(s string, sub string) string {
	val, _ := (regexp2.MustCompile(sub, 0)).FindStringMatch(s)
	if (val != nil) { return val.String() }
	return ""
}

func isValid_part1(s string) bool {
	return  regexMatch_part1(s, "ecl:.") &&
			regexMatch_part1(s, "pid:.") &&
			regexMatch_part1(s, "eyr:.") &&
			regexMatch_part1(s, "hcl:.") &&
			regexMatch_part1(s, "byr:.") &&
			regexMatch_part1(s, "iyr:.") &&
			regexMatch_part1(s, "hgt:.")
}

func regexMatch_part1(s string, sub string) bool {
	regex := regexp.MustCompile(sub)
	matches := regex.FindAllStringIndex(s, -1)

	return len(matches) >= 1
}