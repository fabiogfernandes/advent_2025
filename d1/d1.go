package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readMovesFile(filepath string) []string {
	//Read moves from file
	dat, err := os.ReadFile(filepath)
	check(err)
	moves := strings.Split(string(dat), "\n")
	return moves
}

func move(dialPointerStart int, moveValue int) (int, int, int) {
	timesPointingZero := 0
	timesPassedZero := 0
	dialPointerEnd := 0
	tempMoveValue := dialPointerStart + moveValue

	if tempMoveValue < 0 {
		dialPointerEnd = tempMoveValue % 100
		if dialPointerEnd < 0 {
			dialPointerEnd = 100 + dialPointerEnd
		}
		timesPassedZero = tempMoveValue * -1 / 100
		if dialPointerStart > 0 {
			timesPassedZero += 1
		}
	} else {
		dialPointerEnd = tempMoveValue % 100
		timesPassedZero = tempMoveValue / 100
		if dialPointerEnd == 0 && tempMoveValue == 0 {
			timesPassedZero += 1
		}
	}

	if dialPointerEnd == 0 {
		timesPointingZero += 1
	}
	return dialPointerEnd, timesPointingZero, timesPassedZero
}

func main() {
	moves := readMovesFile("puzzle_input2.txt")
	timesPointingZero, timesPassedZero, t0, tp0 := 0, 0, 0, 0

	dialPointer := 50 //start point
	fmt.Println("Dial pointing: ", dialPointer)
	for i, m := range moves {

		moveValue, convErr := strconv.Atoi(string([]rune(m)[1:]))
		if convErr != nil {
			fmt.Println("File malformed at line", i)
			return
		}
		if string([]rune(m)[0]) == "L" {
			moveValue = moveValue * -1
		}
		dialPointer, t0, tp0 = move(dialPointer, moveValue)
		timesPointingZero += t0
		timesPassedZero += tp0
	}
	fmt.Println("Times pointer left in zero: ", timesPointingZero)
	fmt.Println("Times pointer passed in zero: ", timesPassedZero)
	// Times pointer left in zero:  1055
	// Times pointer passed in zero:  6386
}
