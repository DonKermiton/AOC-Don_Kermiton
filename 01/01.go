package s01

import (
	share "donkermiton.main/share"
	"fmt"
	"strconv"
)

var wordsDictionary = map[string]string{
	"one":   "",
	"two":   "",
	"three": "",
	"four":  "",
	"five":  "",
	"six":   "",
	"seven": "",
	"eight": "",
	"nine":  "",
}

type Amplitude struct {
	currentNumber int
	wasAssign     bool
}

func Run() {
	scanner, err := share.ReadFile("./01/data.txt")

	if nil != err {
		return
	}

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		leftValue, rightValue := readText(line)
		convertedString := fmt.Sprintf("%d%d", leftValue, rightValue)
		if number, err := strconv.Atoi(convertedString); err == nil {
			sum += number
		}

	}

	fmt.Println(sum)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func readText(line string) (int, int) {
	firstValue := Amplitude{currentNumber: -1, wasAssign: false}
	lastValue := Amplitude{currentNumber: -1, wasAssign: false}

	for i := 0; i < len(line); i++ {
		isNumber := isNumeric(string(line[i]))
		if isNumber {
			clearAllValuesInDictionary()

			parsedNumber := parseToInt(line[i])

			if !firstValue.wasAssign {
				firstValue.currentNumber = parsedNumber
				firstValue.wasAssign = true
			} else if firstValue.wasAssign {
				lastValue.currentNumber = parsedNumber
				lastValue.wasAssign = true
			}
		} else {
			currentChar := rune(line[i])
			result := iterateThroughMap(currentChar)

			if result != -1 {
				fmt.Println(result)
				if !firstValue.wasAssign {
					firstValue.currentNumber = result
					firstValue.wasAssign = true
				} else if firstValue.wasAssign {
					lastValue.currentNumber = result
					lastValue.wasAssign = true
				}
			}

		}
	}

	if lastValue.wasAssign == false {
		lastValue.currentNumber = firstValue.currentNumber
	}

	return firstValue.currentNumber, lastValue.currentNumber
}

func iterateThroughMap(char rune) int {
	index := 0
	for key := range wordsDictionary {
		dictionaryValue := wordsDictionary[key]
		charSuits := checkCharSuits(key, dictionaryValue, char)

		if charSuits {
			wordsDictionary[key] = fmt.Sprintf("%s%s", wordsDictionary[key], string(char))

			if len(wordsDictionary[key]) == len(key) {
				clearSelectedKeyFromDictionary(key)
				return index
			}

		} else {
			clearSelectedKeyFromDictionary(key)
		}

		index++
	}
	return -1
}

func checkCharSuits(key string, value string, char rune) bool {
	valueLength := len(value)
	return rune(key[valueLength]) == char
}

func parseToInt(char uint8) int {
	parsedValue, _ := strconv.Atoi(string(char))

	return parsedValue
}

func clearSelectedKeyFromDictionary(key string) {
	wordsDictionary[key] = ""
}

func clearAllValuesInDictionary() {
	for key := range wordsDictionary {
		clearSelectedKeyFromDictionary(key)
	}
}
