package _2

import (
	"fmt"
	"strconv"
	"unicode"

	"donkermiton.main/share"
)

const (
	green = 13
	blue  = 14
	red   = 12
)

var textMap = map[string]string{
	"green": "",
	"blue":  "",
	"red":   "",
}

func Run() {
	scanner, err := share.ReadFile("./02/data.txt")
	if err != nil {
		return
	}

	index := 0
	resultArr := []int{}
	for scanner.Scan() {
		result := readLine(scanner.Text(), index)

		if result == true {
			fmt.Println(scanner.Text(), index+1)
			resultArr = append(resultArr, index+1)
		}

		index++
	}

	sum := 0
	for _, value := range resultArr {
		sum += value
	}

	fmt.Println(sum)
}

func readLine(line string, index int) bool {
	// skip "GAME {index}: "
	indexDigits := len(fmt.Sprintf("%d", index+1))
	skipEntryRule := 5 + indexDigits + 1

	numberOfCubes := ""
	for i := skipEntryRule; i < len(line); i++ {

		// 32 == empty string
		if line[i] == 32 {
			continue
		}

		// 44 == comma ;; 59 == semicolon
		if line[i] == 44 || line[i] == 59 {
			numberOfCubes = ""
		}

		char := rune(line[i])

		if unicode.IsDigit(char) {
			convertedString := string(char)
			convertedInt, _ := strconv.ParseInt(convertedString, 10, 8)

			numberOfCubes = fmt.Sprintf("%s%d", numberOfCubes, convertedInt)
			clearAllValuesInDictionary()
		} else {
			word := readTextInLine(char)

			if word != "" {
				convertedInt, _ := strconv.ParseInt(numberOfCubes, 10, 8)
				isSetValid := checkNumbersAreValid(convertedInt, word)

				if isSetValid == false {
					return false
				}
			}
		}
	}
	return true
}

func readTextInLine(char rune) string {
	for key := range textMap {
		valueLength := len(textMap[key])

		if rune(key[valueLength]) == char {
			textMap[key] = fmt.Sprintf("%s%s", textMap[key], string(char))
			if len(key) == len(textMap[key]) {
				clearDictionaryValuyByKey(key)
				return key
			}
		}
	}

	return ""
}

func clearDictionaryValuyByKey(key string) {
	textMap[key] = ""
}

func clearAllValuesInDictionary() {
	for key := range textMap {
		textMap[key] = ""
	}
}

func checkNumbersAreValid(numb int64, color string) bool {
	switch color {
	case "green":
		return numb <= green
	case "red":
		return numb <= red
	case "blue":
		return numb <= blue
	}

	return false
}
