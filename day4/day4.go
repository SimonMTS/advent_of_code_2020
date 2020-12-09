package day4

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"regexp"
	"regexp2"
	"strconv"
)

// Run is the pseudo main for this day
func Run() {
	inputStrings := bp.Ntos(bp.GetNN("day4/day4_input.txt"))

	validP1 := 0
	validP2 := 0
	for _, s := range inputStrings {
		if isValid(s) {
			validP2++
		}
		if isValidPart1(s) {
			validP1++
		}
	}

	fmt.Println("\tpart1:", validP1)
	fmt.Println("\tpart2:", validP2)
}

func isValid(s string) bool {
	byr, err := strconv.Atoi(regexMatch(s, "(?<=byr:)\\d{4}"))
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(regexMatch(s, "(?<=iyr:)\\d{4}"))
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(regexMatch(s, "(?<=eyr:)\\d{4}"))
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	hgtSuffix := regexMatch(s, "(?<=hgt:\\d{2,3})\\D\\D")
	hgt, err := strconv.Atoi(regexMatch(s, "(?<=hgt:)\\d{2,3}"))
	if err != nil ||
		(hgtSuffix == "cm" && (hgt < 150 || hgt > 193)) ||
		(hgtSuffix == "in" && (hgt < 59 || hgt > 76)) ||
		(hgtSuffix != "cm" && hgtSuffix != "in") {
		return false
	}

	hcl := regexMatch(s, "hcl:#[0-9a-fA-F]{6}")
	if len(hcl) != 11 {
		return false
	}

	ecl := regexMatch(s, "ecl:(amb|blu|brn|gry|grn|hzl|oth)")
	if len(ecl) != 7 {
		return false
	}

	pid, err := strconv.Atoi(regexMatch(s, "(?<=pid:)\\d{9}(?=\\D|$)"))
	if err != nil || pid < 0 {
		return false
	}

	return true
}

func regexMatch(s string, sub string) string {
	val, _ := (regexp2.MustCompile(sub, 0)).FindStringMatch(s)
	if val != nil {
		return val.String()
	}
	return ""
}

func isValidPart1(s string) bool {
	return regexMatchPart1(s, "ecl:.") &&
		regexMatchPart1(s, "pid:.") &&
		regexMatchPart1(s, "eyr:.") &&
		regexMatchPart1(s, "hcl:.") &&
		regexMatchPart1(s, "byr:.") &&
		regexMatchPart1(s, "iyr:.") &&
		regexMatchPart1(s, "hgt:.")
}

func regexMatchPart1(s string, sub string) bool {
	regex := regexp.MustCompile(sub)
	matches := regex.FindAllStringIndex(s, -1)

	return len(matches) >= 1
}
