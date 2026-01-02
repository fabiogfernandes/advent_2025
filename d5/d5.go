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

	var fresh_ingredients [][]int
	var ingredients []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		rangeStart, rangeEnd := parseRange(scanner.Text())
		fresh_ingredients = append(fresh_ingredients, []int{rangeStart, rangeEnd})
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

	return fresh_ingredients, ingredients, nil
}

func main() {
	fresh_ingredients, ingredients, err := readIngredientsDBFile("puzzle_input2.txt")
	if err != nil {
		panic(err)
	}

	validIngredients := 0
	for _, ingredient := range ingredients {
		isValid := false
		for _, fresh_ingredient := range fresh_ingredients {
			if ingredient >= fresh_ingredient[0] && ingredient <= fresh_ingredient[1] {
				isValid = true
				break
			}
		}
		if isValid {
			validIngredients += 1
		}
	}

	fmt.Println("Sum of valid ingredients:", validIngredients)
}
