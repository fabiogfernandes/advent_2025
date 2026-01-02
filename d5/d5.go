package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Part 2

	fmt.Println("Sum of valid ingredients:", validIngredients)
}
