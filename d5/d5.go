package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseRange(rangeString string) (int, int) {
	values := strings.Split(string(rangeString), "-")
	rangeStart, convErr1 := strconv.Atoi(values[0])
	rangeEnd, convErr2 := strconv.Atoi(values[1])
	if convErr1 != nil || convErr2 != nil {
		panic("Ranges malformed")
	}
	return rangeStart, rangeEnd
}

func readIngredientsDBFile(filename string) ([][]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var freshIngredientRanges [][]int
	var ingredients []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		rangeStart, rangeEnd := parseRange(scanner.Text())
		freshIngredientRanges = append(freshIngredientRanges, []int{rangeStart, rangeEnd})
	}
	for scanner.Scan() {
		ingredient, convErr := strconv.Atoi(scanner.Text())
		if convErr != nil {
			return nil, nil, convErr
		}
		ingredients = append(ingredients, ingredient)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return freshIngredientRanges, ingredients, nil
}

type RangePoint struct {
	value   int
	isStart bool
}

func main() {
	freshIngredientRanges, ingredients, err := readIngredientsDBFile("puzzle_input2.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	validIngredients := 0
	for _, ingredient := range ingredients {
		isValid := false
		for _, freshIngredient := range freshIngredientRanges {
			if ingredient >= freshIngredient[0] && ingredient <= freshIngredient[1] {
				isValid = true
				break
			}
		}
		if isValid {
			validIngredients += 1
		}
	}
	fmt.Println("Sum of valid ingredients:", validIngredients)

	// Part 2
	// Computation takes too long
	// freshIngredients := make(map[int]int)
	// fmt.Println("Number of ranges:", len(freshIngredientRanges))

	// for _, freshIngredient := range freshIngredientRanges {
	// 	fmt.Println("Range start:", freshIngredient[0], "Range end:", freshIngredient[1])
	// 	for i := freshIngredient[0]; i <= freshIngredient[1]; i++ {
	// 		fmt.Println("i:", i)
	// 		freshIngredients[i] = i
	// 	}
	// }

	mergedRanges := [][]int{}
	rangePoints := []RangePoint{}
	for _, freshIngredientRange := range freshIngredientRanges {
		rangePoints = append(rangePoints, RangePoint{value: freshIngredientRange[0], isStart: true})
		rangePoints = append(rangePoints, RangePoint{value: freshIngredientRange[1], isStart: false})
	}

	sort.Slice(rangePoints, func(i, j int) bool {
		if rangePoints[i].value == rangePoints[j].value {
			return rangePoints[i].isStart
		} else {
			return rangePoints[i].value < rangePoints[j].value
		}
	})

	fmt.Println("rangePoints:", rangePoints)
	// sort.Slice(rangePoints, func(i, j) bool {
	// 	return rangePoints[i].value < rangePoints[j].value
	// })

	rangeCounter := 0
	currentRange := []int{}
	for _, rp := range rangePoints {
		if rp.isStart {
			if rangeCounter == 0 {
				currentRange = append(currentRange, rp.value)
			}
			rangeCounter++
		} else {
			rangeCounter--
			if rangeCounter == 0 {
				currentRange = append(currentRange, rp.value)
				mergedRanges = append(mergedRanges, currentRange)
				currentRange = []int{}
			}
		}
	}
	freshIngredients := 0
	for _, mr := range mergedRanges {
		freshIngredients += (mr[1] - mr[0] + 1)
	}

	fmt.Println("freshIngredientRanges:", freshIngredientRanges)
	fmt.Println()
	fmt.Println("merged ranges:", mergedRanges)
	fmt.Println()
	fmt.Println("Number of fresh ingredients:", freshIngredients)
}
