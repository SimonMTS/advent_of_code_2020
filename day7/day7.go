package day7

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func Run() {
	input_strings := bp.GetN("day7/day7_input.txt")
	
	bags := make(map[string][]string)
	for _, s := range input_strings {
		split := (regexp.MustCompile("(( bags contain)|( bag(s|),)|( bag(s|)\\.))")).Split(s, -1)
		bags[split[0]] = split[1:]
	}

	sum := 0
	for key, _ := range bags {
		if containsShinyGold(key, &bags) { sum++ }
	}

	fmt.Print("\tpart 1: "); fmt.Println(sum);
	fmt.Print("\tpart 2: "); fmt.Println(countSubBags("shiny gold", &bags));
}

func containsShinyGold(key string, bags *map[string][]string) bool {
	sub_bags := (*bags)[key]

	for _, val := range sub_bags {
		split := (regexp.MustCompile("\\d+\\s+")).Split(val, -1)[1:]

		if len(split) == 0 { return false }
		if (split[0] == "shiny gold") { return true }
		if containsShinyGold(split[0], bags) { return true }
	}

	return false
}

func countSubBags(key string, bags *map[string][]string) int {
	sub_bags := (*bags)[key]

	count := 0
	for _, val := range sub_bags {
		split := (regexp.MustCompile("\\d+\\s+")).Split(val, -1)

		if len(split) < 2 { continue }
		sub_count, _ := strconv.Atoi(strings.Split(val," ")[1])
		sub_count += countSubBags(split[1], bags) * sub_count
		count += sub_count
	}

	return count
}
