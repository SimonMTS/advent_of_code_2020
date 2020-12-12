package day7

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Run is the pseudo main for this day
func Run() {
	inputStrings := bp.GetN("day7/day7_input.txt")

	bags := make(map[string][]string)
	for _, s := range inputStrings {
		split := (regexp.MustCompile("(( bags contain)|( bag(s|),)|( bag(s|)\\.))")).Split(s, -1)
		bags[split[0]] = split[1:]
	}

	sum := 0
	for key := range bags {
		if containsShinyGold(key, &bags) {
			sum++
		}
	}

	fmt.Println("\tpart 1:", sum)
	fmt.Println("\tpart 2:", countSubBags("shiny gold", &bags))
}

var cache = make(map[string]bool)

func containsShinyGold(key string, bags *map[string][]string) bool {
	if val, ok := cache[key]; ok {
		return val
	}

	subBags := (*bags)[key]

	for _, val := range subBags {
		split := (regexp.MustCompile("\\d+\\s+")).Split(val, -1)[1:]

		if len(split) == 0 {
			cache[key] = false
			return false
		}
		if split[0] == "shiny gold" {
			cache[key] = true
			return true
		}
		if containsShinyGold(split[0], bags) {
			cache[key] = true
			return true
		}
	}

	cache[key] = false
	return false
}

func countSubBags(key string, bags *map[string][]string) int {
	subBags := (*bags)[key]

	count := 0
	for _, val := range subBags {
		split := (regexp.MustCompile("\\d+\\s+")).Split(val, -1)

		if len(split) < 2 {
			continue
		}
		subCount, _ := strconv.Atoi(strings.Split(val, " ")[1])
		subCount += countSubBags(split[1], bags) * subCount
		count += subCount
	}

	return count
}
