package _2v2

import (
	"fmt"
	"strconv"
	"unicode"

	"donkermiton.main/share"
)

var textMap = map[string]string{
	"green": "",
	"blue":  "",
	"red":   "",
}

type HighestNumbers struct {
	green int64
	red   int64
	blue  int64
}

func Run() {
	scanner, err := share.ReadFile("./02/data.txt")
	if err != nil {
		return
	}

	index := 0
	var sum int64 = 0
	for scanner.Scan() {
		result := readLine(scanner.Text(), index)

		multiplied := result.green * result.blue * result.red

		sum += multiplied

		index++
	}

	fmt.Println(sum)
}

func readLine(line string, index int) HighestNumbers {
	// skip "GAME {index}: "
	indexDigits := len(fmt.Sprintf("%d", index+1))
	skipEntryRule := 5 + indexDigits + 1

	h := HighestNumbers{
		green: 0,
		blue:  0,
		red:   0,
	}

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
				h.checkNumbersAreValid(convertedInt, word)
				//isSetValid := checkNumbersAreValid(convertedInt, word)
			}
		}
	}
	return h
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

func (h *HighestNumbers) checkNumbersAreValid(numb int64, color string) {
	switch color {
	case "green":
		if h.green < numb {
			h.green = numb
		}
	case "red":
		if h.red < numb {
			h.red = numb
		}
	case "blue":
		if h.blue < numb {
			h.blue = numb
		}
	}
}
