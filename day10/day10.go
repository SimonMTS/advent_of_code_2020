package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"sort"
)

// Run is the pseudo main for this day
func Run() {
	inputInts := bp.Atoi(bp.GetN("day10/day10_input.txt"))
	sort.Ints(inputInts)

	oneJoltDiff, threeJoltDiff := 0, 0
	for i := 0; i <= len(inputInts); i++ {
		curr := 0
		if i == 0 {
			curr = 0
		} else {
			curr = inputInts[i-1]
		}

		next := 0
		if len(inputInts) == i {
			next = inputInts[i-1] + 3
		} else {
			next = inputInts[i]
		}

		if next-curr == 1 {
			oneJoltDiff++
		} else if next-curr == 3 {
			threeJoltDiff++
		}
	}

	fmt.Println("\tpart 1:", oneJoltDiff*threeJoltDiff)

	inputInts = append([]int{0}, inputInts...)
	inputInts = append(inputInts, inputInts[len(inputInts)-1]+3)

	indices := make([]int, 1)
	for i := 1; i < len(inputInts); i++ {
		if inputInts[i-1] == inputInts[i]-3 {
			indices = append(indices, i)
		}
	}

	res := 1
	for i := 1; i < len(indices); i++ {
		paths, prevIndex, currIndex := 0, indices[i-1], indices[i]
		next(prevIndex, &inputInts, &paths, currIndex)
		res *= paths
	}

	fmt.Println("\tpart 2:", res)
}

func next(index int, list *[]int, paths *int, goal int) int {
	if index == goal {
		(*paths)++
		return 1
	}

	newPaths := 0
	val := (*list)[index]

	nPlus1 := 0
	if index+1 < len(*list) {
		nPlus1 = (*list)[index+1]
	} else {
		nPlus1 = val + 5
	}

	nPlus2 := 0
	if index+2 < len(*list) {
		nPlus2 = (*list)[index+2]
	} else {
		nPlus2 = val + 5
	}

	nPlus3 := 0
	if index+3 < len(*list) {
		nPlus3 = (*list)[index+3]
	} else {
		nPlus3 = val + 5
	}

	if nPlus1-val <= 3 {
		newPaths += 1 + next(index+1, list, paths, goal)
	}

	if nPlus2-val <= 3 {
		newPaths += 1 + next(index+2, list, paths, goal)
	}

	if nPlus3-val <= 3 {
		newPaths += 1 + next(index+3, list, paths, goal)
	}

	return newPaths
}
