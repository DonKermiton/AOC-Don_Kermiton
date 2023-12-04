package s01

import (
	share "donkermiton.main/share"
	"fmt"
	"strconv"
)

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

func readText(line string) (int, int) {
	var leftPointer int = -1
	var rightPointer int = -1

	for i := 0; i < len(line); i++ {
		leftValue, err := parseToInt(line[i])

		if err == nil && leftPointer == -1 {
			leftPointer = leftValue
		}

		rightValue, err := parseToInt(line[len(line)-i-1])

		if err == nil && rightPointer == -1 {
			rightPointer = rightValue
		}

		if leftPointer != -1 && rightPointer != -1 {
			break
		}
	}

	return leftPointer, rightPointer
}

func parseToInt(char uint8) (int, error) {
	parsedValue, err := strconv.Atoi(string(char))

	if err != nil {
		return -1, err
	}

	return parsedValue, nil
}
