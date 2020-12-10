package main

import (
	day1 "advent_of_code_2020/day1"
	day10 "advent_of_code_2020/day10"
	day2 "advent_of_code_2020/day2"
	day3 "advent_of_code_2020/day3"
	day4 "advent_of_code_2020/day4"
	day5 "advent_of_code_2020/day5"
	day6 "advent_of_code_2020/day6"
	day7 "advent_of_code_2020/day7"
	day8 "advent_of_code_2020/day8"
	day9 "advent_of_code_2020/day9"
	"strings"

	// day11 "advent_of_code_2020/day11"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day := 0
	if len(os.Args[1:]) == 1 {
		if len(os.Args[1]) > 3 {
			a := strings.Split(os.Args[1], "\\")
			day, _ = strconv.Atoi(string(a[len(a)-1][3:]))
		} else {
			day, _ = strconv.Atoi(os.Args[1])
		}
	}

	if day == 0 || day == 1 {
		fmt.Println("\n    day1")
		day1.Run()
	}

	if day == 0 || day == 2 {
		fmt.Println("\n    day2")
		day2.Run()
	}

	if day == 0 || day == 3 {
		fmt.Println("\n    day3")
		day3.Run()
	}

	if day == 0 || day == 4 {
		fmt.Println("\n    day4")
		day4.Run()
	}

	if day == 0 || day == 5 {
		fmt.Println("\n    day5")
		day5.Run()
	}

	if day == 0 || day == 6 {
		fmt.Println("\n    day6")
		day6.Run()
	}

	if day == 0 || day == 7 {
		fmt.Println("\n    day7")
		day7.Run()
	}

	if day == 0 || day == 8 {
		fmt.Println("\n    day8")
		day8.Run()
	}

	if day == 0 || day == 9 {
		fmt.Println("\n    day9")
		day9.Run()
	}

	if day == 0 || day == 10 {
		fmt.Println("\n    day10")
		day10.Run()
	}

	if day == 0 || day == 11 {
		fmt.Println("\n    day11")
		// day11.Run()
	}

	fmt.Println("")
}
