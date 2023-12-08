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
	for scanner.Scan() {

		readLine(scanner.Text(), index)
		break
		index++
	}
}

func readLine(line string, index int) {
	// skip "GAME {index}: "
	indexDigits := len(fmt.Sprintf("%d", index+1))
	skipEntryRule := 5 + indexDigits + 1

	var numberOfCubes []int = []int{}
	for i := skipEntryRule; i < len(line); i++ {

		// 32 == empty string
		if line[i] == 32 {
			continue
		}

		// 44 == comma
		if line[i] == 44 || line[i] == 59 {
			numberOfCubes = []int{}
		}

		char := rune(line[i])

		if unicode.IsDigit(char) {
			convertedString := string(char)
			convertedInt, _ := strconv.ParseInt(convertedString, 10, 8)

			numberOfCubes = append(numberOfCubes, int(convertedInt))
			clearAllValuesInDictionary()
		} else {
			word := readTextInLine(char)

			if word != "" {
				fmt.Println(numberOfCubes, word)
			}
		}
	}
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
