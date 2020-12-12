package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
)

// Run is the pseudo main for this day
func Run() {
	input := bp.GetN("day11/day11_input.txt")

	seats1 := make([][]string, len(input))
	seats2 := make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		seats1[i] = make([]string, len(input[i]))
		seats2[i] = make([]string, len(input[i]))
	}

	for col, l := range input {
		for row, c := range l {
			seats1[col][row] = string(c)
			seats2[col][row] = string(c)
		}
	}

	fmt.Println("\tpart 1:", doSim(4, seats1, input))
	fmt.Println("\tpart 2:", doSim(5, seats2, input))

}

func doSim(mode int, seats [][]string, input []string) int {
	newSeats := make([][]string, len(input))
	for i := range seats {
		newSeats[i] = make([]string, len(seats[i]))
		copy(newSeats[i], seats[i])
	}

	for true {
		if !applyCicle(mode, &seats, &newSeats) {
			break
		}

		for i := range newSeats {
			// fmt.Println(newSeats[i])
			seats[i] = make([]string, len(newSeats[i]))
			copy(seats[i], newSeats[i])
		}
		// fmt.Println("")
	}

	occupied := 0
	for _, l := range seats {
		// fmt.Println(l)
		for _, c := range l {
			if c == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func applyCicle(mode int, seats, newSeats *[][]string) bool {
	changes := 0
	for col, l := range *seats {
		for row := range l {
			changes += updateSeat(mode, col, row, seats, newSeats)
		}
	}
	return changes != 0
}

func updateSeat(mode, col, row int, seats, newSeats *[][]string) int {
	// 1 2 3
	// 8 x 4
	// 7 6 5
	occupied := 0
	_ = getSeat(1, col, row, seats, &occupied, mode == 4)
	_ = getSeat(2, col, row, seats, &occupied, mode == 4)
	_ = getSeat(3, col, row, seats, &occupied, mode == 4)
	_ = getSeat(4, col, row, seats, &occupied, mode == 4)
	_ = getSeat(5, col, row, seats, &occupied, mode == 4)
	_ = getSeat(6, col, row, seats, &occupied, mode == 4)
	_ = getSeat(7, col, row, seats, &occupied, mode == 4)
	_ = getSeat(8, col, row, seats, &occupied, mode == 4)

	change := 0
	if (*seats)[col][row] == "L" && occupied == 0 {
		(*newSeats)[col][row] = "#"
		change++
	} else if (*seats)[col][row] == "#" && occupied >= mode {
		(*newSeats)[col][row] = "L"
		change++
	}

	return change
}

func getSeat(dir, col, row int, seats *[][]string, occupied *int, justFirst bool) string {
	// 1 2 3
	// 8 x 4
	// 7 6 5
	realRow, realCol := row, col

	if dir == 1 || dir == 2 || dir == 3 {
		realCol--
	}
	if dir == 5 || dir == 6 || dir == 7 {
		realCol++
	}

	if dir == 1 || dir == 8 || dir == 7 {
		realRow--
	}
	if dir == 3 || dir == 4 || dir == 5 {
		realRow++
	}

	seat, final := getSingleSeat(realCol, realRow, seats)
	stepCol, stepRow := realCol-col, realRow-row

	for !(justFirst || seat == "#" || seat == "L" || final) {
		realCol += stepCol
		realRow += stepRow

		seat, final = getSingleSeat(realCol, realRow, seats)
	}

	if seat == "#" {
		*occupied++
	}
	return seat
}

func getSingleSeat(col, row int, seats *[][]string) (string, bool) {
	if !(col >= 0) {
		return ".", true
	}
	if !(col < len((*seats))) {
		return ".", true
	}

	if !(row >= 0) {
		return ".", true
	}
	if !(row < len((*seats)[col])) {
		return ".", true
	}

	seat := (*seats)[col][row]
	return seat, false
}
