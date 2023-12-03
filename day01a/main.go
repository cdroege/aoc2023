package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getCalibrationValueFromLine(line string) int {
	numbers := []int{}

	for _, char := range line {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char-'0'))
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
