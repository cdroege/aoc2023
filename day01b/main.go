package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var writtenNumberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValueFromLine(line string) int {
	numbers := []int{}

	for index, char := range line {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char-'0'))
		} else {
			// Brute force the solution by just looking ahead
			// if any key from writtenNumberMap matches the
			// next characters
			for key := range writtenNumberMap {
				if index+len(key) > len(line) {
					// Skip key due to EOL
					continue
				}

				maybeMatch := line[index : index+len(key)]

				if key == maybeMatch {
					numbers = append(numbers, writtenNumberMap[key])
				}
			}
		}
	}

	if len(numbers) == 0 {
		return 0
	}

	return numbers[0]*10 + numbers[len(numbers)-1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	calibrationLines := strings.Split(string(input), "\n")
	for _, line := range calibrationLines {
		sum += getCalibrationValueFromLine(line)
	}
	fmt.Println("Calibration value:", sum)
}
