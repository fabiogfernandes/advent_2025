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

func main() {
	//Read moves from file
	dat, err := os.ReadFile("puzzle_input2.txt")
	check(err)
	moves := strings.Split(string(dat), "\n")
	timesPointingZero := 0
	dialPointer := 50 //start point

	for i, m := range moves {
		moveValue, convErr := strconv.Atoi(string([]rune(m)[1:]))
		if convErr != nil {
			fmt.Println("File malformed at line", i)
			return
		}
		if string([]rune(m)[0]) == "L" {
			moveValue = moveValue * -1
		}
		dialPointer = (100 + dialPointer + moveValue) % 100
		if dialPointer == 0 {
			timesPointingZero++
		}
	}
	fmt.Println("Times pointer left in zero: ", timesPointingZero)
}
