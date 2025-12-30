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

func findLargest(numString string) (int, int) {
	firstIndexLargest := -1
	largest := 0
	for i, n := range numString {
		batteryValue := int(n) - int('0')
		// batteryValue, convErr := strconv.Atoi(n)
		if batteryValue > largest {
			largest = batteryValue
			firstIndexLargest = i
		}
	}
	return largest, firstIndexLargest
}

func largestJoltageNBatteries(numString string, numBatteries int) int {
	indexLargest, largest, tempBank := -1, 0, numString
	var activeBatteries strings.Builder
	for batteriesLeft := numBatteries; batteriesLeft > 0; batteriesLeft-- {
		// fmt.Println("SubString: ", tempBank[:len(tempBank)-batteriesLeft+1])
		largest, indexLargest = findLargest(tempBank[:len(tempBank)-batteriesLeft+1])
		activeBatteries.WriteString(strconv.Itoa(largest))
		tempBank = tempBank[indexLargest+1:]
	}
	retValue, convErr := strconv.Atoi(activeBatteries.String())
	if convErr != nil {
		panic(convErr)
	}
	return retValue
}

func readBatteryBanksFile(filepath string) []string {
	//Read battery banks line from file
	dat, err := os.ReadFile(filepath)
	check(err)
	banks := strings.Split(string(dat), "\n")
	return banks
}

func main() {

	banks := readBatteryBanksFile("puzzle_input2.txt")
	firstIndexLargest, firstBattery, secondBattery, joltage, totalJoltage := -1, 0, 0, 0, 0
	totalJoltage2, joltage2 := 0, 0
	for _, b := range banks {
		// Part 1
		// fmt.Println("Bank: ", b)
		firstBattery, firstIndexLargest = findLargest(b[0 : len(b)-1])
		// fmt.Printf("FB: %i - FIL: %i", firstBattery, firstIndexLargest)
		secondBattery, _ = findLargest(b[firstIndexLargest+1:])
		joltage = firstBattery*10 + secondBattery
		// fmt.Println("Bank: ", b)
		// fmt.Println("Joltage:", joltage)
		totalJoltage += joltage
		// Part 2
		joltage2 = largestJoltageNBatteries(b, 12)
		totalJoltage2 += joltage2

	}
	fmt.Println("Total joltage: ", totalJoltage)
	fmt.Println("Total joltage2: ", totalJoltage2)

	// fmt.Println(largestJoltageNBatteries("811111111111119", 12))

}
