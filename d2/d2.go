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

func readRangesFile(filepath string) []string {
	//Read ranges from file
	dat, err := os.ReadFile(filepath)
	check(err)
	ranges := strings.Split(string(dat), ",")
	return ranges
}

func parseRange(rangeString string) (int, int) {
	values := strings.Split(string(rangeString), "-")
	rangeStart, convErr1 := strconv.Atoi(values[0])
	rangeEnd, convErr2 := strconv.Atoi(values[1])
	if convErr1 != nil || convErr2 != nil {
		panic("Ranges malformed")
	}
	return rangeStart, rangeEnd
}

func main() {
	ranges := readRangesFile("puzzle_input2.txt")
	// numberString, firstHalf, secondHalf := "", "", ""
	numberString := ""
	middleIndex := 0
	invalidsSum := 0
	// convErr1, convErr2 = nil, nil

	for _, r := range ranges {
		rangeStart, rangeEnd := parseRange(r)
		// fmt.Printf("%d - %d \n", rangeStart, rangeEnd)
		for i := rangeStart; i <= rangeEnd; i++ {
			numberString = strconv.Itoa(i)
			if len(numberString)%2 == 0 {
				middleIndex = len(numberString) / 2
				if numberString[0:middleIndex] == numberString[middleIndex:] {
					// fmt.Println("Invalid: ", numberString)
					invalidsSum += i
				}
			}
		}
	}
	fmt.Println("Invalid IDs sum: ", invalidsSum)
}
