package _4

import (
	"testing"
)

func checkSets(t *testing.T, winningNumbers []int, correctWinningNumbers []int) {
	if len(winningNumbers) != len(correctWinningNumbers) {
		t.Error("First set has wrong length diff")
	}

	for i, _ := range winningNumbers {
		if winningNumbers[i] != correctWinningNumbers[i] {
			t.Errorf("Different Numbers;; Expected %d, got %d", correctWinningNumbers[i], winningNumbers[i])
		}
	}
}

func checkResultSets(t *testing.T, algoResult []int, validResult []int) {
	for _, number := range algoResult {
		if IsWinningNumber(validResult, number) == false {
			t.Error("Missing number in set \t", number)
		}
	}
}

func TestFirstSet(t *testing.T) {
	result, winningNumbers := readLine("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 0)
	correctWinningNumbers := []int{41, 48, 83, 86, 17}
	correctResultNumbers := []int{83, 86, 17, 48}

	checkSets(t, winningNumbers, correctWinningNumbers)
	checkResultSets(t, result, correctResultNumbers)
}

func TestSecondSet(t *testing.T) {
	result, winningNumbers := readLine("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 1)
	correctWinningNumbers := []int{13, 32, 20, 16, 61}
	correctResultNumbers := []int{32, 61}
	checkSets(t, winningNumbers, correctWinningNumbers)

	checkResultSets(t, result, correctResultNumbers)
}

func TestThirdSet(t *testing.T) {
	result, winningNumbers := readLine("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2)
	correctWinningNumbers := []int{1, 21, 53, 59, 44}
	correctResultNumbers := []int{1, 21}
	checkSets(t, winningNumbers, correctWinningNumbers)

	checkResultSets(t, result, correctResultNumbers)
}

func TestFourthSet(t *testing.T) {
	result, winningNumbers := readLine("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 3)
	correctWinningNumbers := []int{41, 92, 73, 84, 69}
	correctResultNumbers := []int{84}
	checkSets(t, winningNumbers, correctWinningNumbers)

	checkResultSets(t, result, correctResultNumbers)
}

func TestFifthSet(t *testing.T) {
	result, winningNumbers := readLine("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 4)
	correctWinningNumbers := []int{87, 83, 26, 28, 32}
	checkSets(t, winningNumbers, correctWinningNumbers)

	if len(result) != 0 {
		t.Error("should be empty array")
	}
}

func TestSixthSet(t *testing.T) {
	result, winningNumbers := readLine("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 5)
	correctWinningNumbers := []int{31, 18, 13, 56, 72}

	checkSets(t, winningNumbers, correctWinningNumbers)
	if len(result) != 0 {
		t.Error("should be empty array")
	}
}

func TestCheckIsWinningNumber(t *testing.T) {
	isWinningNumber := IsWinningNumber([]int{31, 18, 13, 56, 72}, 31)

	if isWinningNumber == false {
		t.Error("Number should be true \t", 13)
	}

	isWinningNumber = IsWinningNumber([]int{31, 18, 13, 56, 72}, 1)

	if isWinningNumber == true {
		t.Error("Number should be false \t", 1)
	}

}
