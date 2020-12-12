package day8

import (
	bp "advent_of_code_2020/boiler_plate"
	"fmt"
	"math"
	"strconv"
)

var dirs = []string{"north", "east", "south", "west"}

// Run is the pseudo main for this day
func Run() {
	input := bp.GetN("day12/day12_input.txt")

	shipX, shipY, shipAxis := 0, 0, 1
	for _, s := range input {
		action := s[:1]
		value, _ := strconv.Atoi(s[1:])
		applyActionPart1(action, value, &shipX, &shipY, &shipAxis)
	}
	fmt.Println("\tpart 1:", shipX+shipY)

	shipX, shipY, waypointX, waypointY := 0, 0, 10, 1
	// fmt.Println(shipX, shipY, waypointX, waypointY)

	for _, s := range input {
		action := s[:1]
		value, _ := strconv.Atoi(s[1:])

		applyActionPart2(action, value, &shipX, &shipY, &waypointX, &waypointY)

		// fmt.Println("[", shipX, shipY, "] [", waypointX, waypointY, "] -", action, value)
	}

	fmt.Println("\tpart 2:", math.Abs(float64(shipX))+math.Abs(float64(shipY)))
}

func applyActionPart1(action string, value int, x, y, shipAxis *int) {
	if action == "F" {
		if dirs[*shipAxis] == "north" {
			*y -= value
		} else if dirs[*shipAxis] == "east" {
			*x += value
		} else if dirs[*shipAxis] == "south" {
			*y += value
		} else if dirs[*shipAxis] == "west" {
			*x -= value
		}
	}

	if action == "N" {
		*y -= value
	} else if action == "E" {
		*x += value
	} else if action == "S" {
		*y += value
	} else if action == "W" {
		*x -= value
	}

	ticks := value / 90
	if action == "R" {
		*shipAxis = (*shipAxis + ticks) % 4
	} else if action == "L" {
		*shipAxis = (*shipAxis + (4 - ticks)) % 4
	}
}

func applyActionPart2(action string, value int, x, y, waypointX, waypointY *int) {
	if action == "F" {
		for i := 0; i < value; i++ {
			*x += *waypointX
			*y += *waypointY
		}
	}

	if action == "N" {
		*waypointY += value
	} else if action == "E" {
		*waypointX += value
	} else if action == "S" {
		*waypointY -= value
	} else if action == "W" {
		*waypointX -= value
	}

	if action == "R" {
		for i := 0; i < value/90; i++ {
			oldX, oldY := *waypointX, *waypointY
			*waypointX = oldY
			*waypointY = -oldX
		}
	} else if action == "L" {
		for i := 0; i < (360 - (value / 90)); i++ {
			oldX, oldY := *waypointX, *waypointY
			*waypointX = oldY
			*waypointY = -oldX
		}
	}
}
