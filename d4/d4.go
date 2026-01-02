package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readRollsDiagramFile(filepath string) [][]rune {
	//Read rolls diagram from file
	dat, err := os.ReadFile(filepath)
	fmt.Println(reflect.TypeOf(dat))
	check(err)
	lines := strings.Split(string(dat), "\n")
	diagram := make([][]rune, len(lines))
	for i, l := range lines {
		diagram[i] = []rune(l)
	}
	return diagram
}

func isRoll(pos rune) bool {
	return pos == '@'
}

func adjacentRolls(posX int, posY int, diagram [][]rune) int {
	lengthX := len(diagram[0])
	lengthY := len(diagram)
	adjacentRolls := 0

	if posX-1 >= 0 {
		if isRoll(diagram[posY][posX-1]) {
			// fmt.Println("X-1 Y")
			adjacentRolls += 1
		}
		if posY-1 >= 0 {
			if isRoll(diagram[posY-1][posX-1]) {
				// fmt.Println("X-1 Y-1")
				adjacentRolls += 1
			}
		}
		if posY+1 < lengthY {
			if isRoll(diagram[posY+1][posX-1]) {
				// fmt.Println("X-1 Y+1")
				adjacentRolls += 1
			}
		}
	}

	if posX+1 < lengthX {
		if isRoll(diagram[posY][posX+1]) {
			// fmt.Println("X+1 Y")
			adjacentRolls += 1
		}
		if posY-1 >= 0 {
			if isRoll(diagram[posY-1][posX+1]) {
				// fmt.Println("X+1 Y-1")
				adjacentRolls += 1
			}
		}
		if posY+1 < lengthY {
			if isRoll(diagram[posY+1][posX+1]) {
				// fmt.Println("X+1 Y+1")
				adjacentRolls += 1
			}
		}
	}

	if posY-1 >= 0 {
		if isRoll(diagram[posY-1][posX]) {
			// fmt.Println("X Y-1")
			adjacentRolls += 1
		}
	}

	if posY+1 < lengthY {
		if isRoll(diagram[posY+1][posX]) {
			// fmt.Println("X Y+1")
			adjacentRolls += 1
		}
	}
	return adjacentRolls
}

func adjacentRollCalc(posX int, posY int, diagram [][]rune) int {
	lengthX := len(diagram[0])
	lengthY := len(diagram)

	if posX >= 0 && posX < lengthX && posY >= 0 && posY < lengthY && isRoll(diagram[posY][posX]) {
		// fmt.Printf("ADJ2: X:%d Y:%d\n", posX, posY)
		return 1
	}
	return 0
}

func adjacentRolls2(posX int, posY int, diagram [][]rune) int {
	ar := 0
	ar += adjacentRollCalc(posX-1, posY, diagram)
	ar += adjacentRollCalc(posX-1, posY-1, diagram)
	ar += adjacentRollCalc(posX-1, posY+1, diagram)
	ar += adjacentRollCalc(posX, posY-1, diagram)
	ar += adjacentRollCalc(posX, posY+1, diagram)
	ar += adjacentRollCalc(posX+1, posY, diagram)
	ar += adjacentRollCalc(posX+1, posY-1, diagram)
	ar += adjacentRollCalc(posX+1, posY+1, diagram)
	return ar
}

func main() {
	diagram := readRollsDiagramFile("puzzle_input2.txt")
	accessableRolls := 0

	// x, y := 9, 0
	// fmt.Printf("Pos %d, %d - value %s \n", x, y, diagram[x][y])
	// fmt.Println("Adj: ", adjacentRolls(x, y, diagram))
	// fmt.Println("Adj2: ", adjacentRolls2(x, y, diagram))

	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] == '@' {
				if adjacentRolls2(j, i, diagram) < 4 {
					accessableRolls += 1
				}
			}
		}
	}
	fmt.Println("Part 1 - Accessable rolls: ", accessableRolls)

	// Part 2

	accessableRollsPositions := [][]int{}
	totalAccessableRolls := 0

	for hasAccessableRolls := true; hasAccessableRolls; {
		for i := 0; i < len(diagram); i++ {
			for j := 0; j < len(diagram[i]); j++ {
				if diagram[i][j] == '@' {
					if adjacentRolls2(j, i, diagram) < 4 {
						accessableRollsPositions = append(accessableRollsPositions, []int{j, i})
					}
				}
			}
		}
		fmt.Println("Accessable rolls: ", len(accessableRollsPositions))
		totalAccessableRolls += len(accessableRollsPositions)
		if len(accessableRollsPositions) == 0 {
			hasAccessableRolls = false
		} else {
			for _, pos := range accessableRollsPositions {
				diagram[pos[1]][pos[0]] = '.'
			}
			accessableRollsPositions = [][]int{}
		}

	}

	fmt.Println("Part 2 - Total Accessable rolls removed: ", totalAccessableRolls)
}
