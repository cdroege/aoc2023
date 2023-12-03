package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cubeSet struct {
	red   int
	green int
	blue  int
}

func isPossibleCubeGame(gameConfiguration cubeSet, cubeSubsets []cubeSet) bool {
	for _, subset := range cubeSubsets {
		if gameConfiguration.red < subset.red {
			return false
		}

		if gameConfiguration.green < subset.green {
			return false
		}

		if gameConfiguration.blue < subset.blue {
			return false
		}
	}

	return true
}

func parseLineToIDandCubeSubset(line string) (int, []cubeSet, error) {
	gameIDStr, cubeSubsetsStr, found := strings.Cut(line, ": ")
	if !found {
		return 0, nil, errors.New("No game description was found in this line")
	}

	gameIDStr, found = strings.CutPrefix(gameIDStr, "Game ")
	if !found {
		return 0, nil, errors.New("Game ID could not be extracted")
	}

	gameID, err := strconv.Atoi(gameIDStr)

	if err != nil {
		return 0, nil, err
	}

	cubeSubsets := []cubeSet{}
	for _, subset := range strings.Split(cubeSubsetsStr, "; ") {
		red := 0
		green := 0
		blue := 0

		for _, cubeStr := range strings.Split(subset, ", ") {
			amountStr, color, found := strings.Cut(cubeStr, " ")
			if !found {
				return 0, nil, errors.New("Color or amount not found in line: " + cubeStr)
			}

			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				return 0, nil, errors.New("Could not parse amount in: " + amountStr)
			}

			switch color {
			case "red":
				red = amount
			case "green":
				green = amount
			case "blue":
				blue = amount
			default:
				return 0, nil, errors.New("Unknown color: " + color)
			}
		}

		cubeSubsets = append(cubeSubsets, cubeSet{red: red, blue: blue, green: green})
	}

	return gameID, cubeSubsets, nil
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sumOfPossibleGameIDs := 0
	gameConfiguration := cubeSet{red: 12, green: 13, blue: 14}
	calibrationLines := strings.Split(string(input), "\n")
	for _, line := range calibrationLines {
		if len(line) == 0 {
			continue
		}

		gameID, cubeSubsets, err := parseLineToIDandCubeSubset(line)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Comparing active configuration: %v\n", gameConfiguration)
		fmt.Printf("With the following cubeSubSets: %v\n", cubeSubsets)

		if isPossibleCubeGame(gameConfiguration, cubeSubsets) {
			fmt.Print("gameID is possible: ", gameID, "\n\n")
			sumOfPossibleGameIDs += gameID
		}
	}
	fmt.Println("Sum of possible game IDs:", sumOfPossibleGameIDs)
}
