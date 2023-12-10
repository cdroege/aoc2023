package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func runeToNumber(input rune) int {
	return int(input - '0')
}

func isSymbol(input rune) bool {
	return !unicode.IsDigit(input) && input != '.'
}

// hasAdjacentSymbol checks if in the upper or lower row or left or right column of a position
// is a symbol
func hasAdjacentSymbol(schematic [][]rune, rowIndex int, colIndex int) bool {
	hasUpperRow := rowIndex > 0
	hasLeftCol := colIndex > 0
	hasLowerRow := rowIndex < len(schematic)-1
	hasRightCol := colIndex < len(schematic[0])-1

	return (hasLeftCol && isSymbol(schematic[rowIndex][colIndex-1])) ||
		(hasRightCol && isSymbol(schematic[rowIndex][colIndex+1])) ||
		(hasLowerRow && isSymbol(schematic[rowIndex+1][colIndex])) ||
		(hasUpperRow && hasRightCol && isSymbol(schematic[rowIndex-1][colIndex+1])) ||
		(hasLowerRow && hasRightCol && isSymbol(schematic[rowIndex+1][colIndex+1])) ||
		(hasLowerRow && hasLeftCol && isSymbol(schematic[rowIndex+1][colIndex-1])) ||
		(hasUpperRow && hasRightCol && isSymbol(schematic[rowIndex-1][colIndex+1])) ||
		(hasUpperRow && hasLeftCol && isSymbol(schematic[rowIndex-1][colIndex-1])) ||
		(hasUpperRow && isSymbol(schematic[rowIndex-1][colIndex]))
}

func findEnginePartNumbers(rawData string) []int {
	numberOfRows := strings.Count(rawData, "\n")
	numberOfCols := (len(rawData) - 1) / numberOfRows

	partNumbers := []int{}

	// Convert the string to a 2D matrix first so that
	// we have easier access to the elements above
	// and below individual runes
	schematic := make([][]rune, numberOfRows)
	for indexRow, line := range strings.Split(rawData, "\n") {
		if len(line) == 0 {
			continue
		}
		schematic[indexRow] = make([]rune, numberOfCols+1)
		for indexCol, character := range line {
			schematic[indexRow][indexCol] = character
		}

		// Add some padding to the right to make the else clause
		// down below simpler
		schematic[indexRow][len(schematic[indexRow])-1] = '.'
	}

	for rowIndex, row := range schematic {
		currentNumber := -1
		isAdjacent := false

		for colIndex, character := range row {
			if unicode.IsDigit(character) {
				number := runeToNumber(character)
				if currentNumber == -1 {
					currentNumber = number
				} else {
					currentNumber = currentNumber*10 + number
				}

				if hasAdjacentSymbol(schematic, rowIndex, colIndex) {
					isAdjacent = true
				}
			} else {
				if isAdjacent && currentNumber != -1 {
					partNumbers = append(partNumbers, currentNumber)
					isAdjacent = false
				}
				currentNumber = -1
			}
		}
	}

	return partNumbers
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Input file is missing")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	partNumbers := findEnginePartNumbers(string(input))

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	fmt.Printf("Sum of part numbers: %d\n", sum)
}
