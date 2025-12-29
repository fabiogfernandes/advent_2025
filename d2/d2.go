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

func isRepeatedLengthSequence(numberString string, length int) bool {
	if len(numberString)%length != 0 {
		return false
	}
	for i := length; i <= len(numberString)-1; i += length {
		// fmt.Println("Compare", numberString[0:length], numberString[i:i+length])
		if numberString[0:length] != numberString[i:i+length] {
			return false
		}
	}
	return true
}

func repeatedSequence(numberString string) bool {
	for repSeqLength := 1; repSeqLength <= len(numberString)/2; repSeqLength++ {
		if isRepeatedLengthSequence(numberString, repSeqLength) {
			return true
		}
	}

	return false
}

func main() {

	ranges := readRangesFile("puzzle_input2.txt")
	numberString := ""
	middleIndex := 0
	invalidsSum, invalidsSum2 := 0, 0

	for _, r := range ranges {
		rangeStart, rangeEnd := parseRange(r)
		// fmt.Printf("%d - %d \n", rangeStart, rangeEnd)
		for n := rangeStart; n <= rangeEnd; n++ {

			numberString = strconv.Itoa(n)

			// Part 1 solution
			if len(numberString)%2 == 0 {
				middleIndex = len(numberString) / 2
				if numberString[0:middleIndex] == numberString[middleIndex:] {
					// fmt.Println("Invalid: ", numberString)
					invalidsSum += n
				}
			}

			// Part 2 solution
			if repeatedSequence(numberString) {
				invalidsSum2 += n
				// fmt.Println("Invalid: ", numberString)
			}
		}
	}
	fmt.Println("Invalid IDs Part 1 sum: ", invalidsSum)
	fmt.Println("Invalid IDs Part 2 sum: ", invalidsSum2)
}
