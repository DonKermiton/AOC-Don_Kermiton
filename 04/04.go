package _4

import (
	"fmt"
	"strconv"

	"donkermiton.main/share"
)

const (
	addingWinningNumbers = "winning"
	addingUserNumbers    = "user"
)

func Run() {
	scanner, err := share.ReadFile("./04/data.txt")

	if err != nil {
		return
	}

	var index int64 = 0

	sum := 0
	for scanner.Scan() {
		result, _ := readLine(scanner.Text(), index)

		points := 0
		for range result {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

		fmt.Println(result)
		sum += points

		index++
	}

	fmt.Println(sum)
}

func readLine(line string, index int64) ([]int, []int) {
	var result []int
	var winningNumbers []int

	mode := addingWinningNumbers
	currentWinningNumbers := ""
	for i := calcLinesToSkip(index); i < len(line); i++ {
		currCharASCII := line[i]
		currChar := string(currCharASCII)

		// ASCII 32 == Empty String
		if currCharASCII == 32 {
			if len(currentWinningNumbers) == 0 {
				continue
			}

			parsedNumber, err := convertStringToNumber(currentWinningNumbers)

			if err != nil {
				fmt.Println("Error while parsing char: \t", currentWinningNumbers, err)
			}

			switch mode {
			case addingWinningNumbers:
				winningNumbers = append(winningNumbers, parsedNumber)
				currentWinningNumbers = ""
			case addingUserNumbers:
				isWinningNumber := IsWinningNumber(winningNumbers, parsedNumber)
				fmt.Println(winningNumbers, parsedNumber)

				if isWinningNumber == true {
					result = append(result, parsedNumber)
				}
				currentWinningNumbers = ""
			}

		}

		if currChar == "|" {
			currentWinningNumbers = ""
			mode = addingUserNumbers
			continue

		}

		if currCharASCII != 32 {
			currentWinningNumbers = fmt.Sprintf("%s%s", currentWinningNumbers, currChar)
		}
	}

	parsedNumber, _ := convertStringToNumber(currentWinningNumbers)

	isWinningNumber := IsWinningNumber(winningNumbers, parsedNumber)
	fmt.Println(winningNumbers, parsedNumber)

	if isWinningNumber == true {
		result = append(result, parsedNumber)
	}
	currentWinningNumbers = ""

	return result, winningNumbers
}

func calcLinesToSkip(index int64) int {
	indexDigits := len(fmt.Sprintf("%d", index+1))
	// skip "Card {index}: "
	return 5 + indexDigits + 2
}

func convertStringToNumber(stringToConvert string) (int, error) {
	return strconv.Atoi(stringToConvert)
}

func IsWinningNumber(winningNumberSets []int, currentNumber int) bool {
	for _, value := range winningNumberSets {
		if value == currentNumber {
			return true
		}
	}

	return false
}
