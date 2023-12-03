package s01

import (
	"testing"
)

func TestFirstTextNumberReader(t *testing.T) {
	left, right := readText("1abc2")

	if left != 1 {
		t.Error("left value, excepted 1 got ", left)
	}

	if right != 2 {
		t.Error("left value, excepted 2 got ", right)
	}
}

func TestSecondTextNumberReader(t *testing.T) {
	left, right := readText("a1b2c3d4e5f")

	if left != 1 {
		t.Error("left value, excepted 1 got ", left)
	}

	if right != 5 {
		t.Error("left value, excepted 5 got ", right)
	}
}

func TestThirdTextNumberReader(t *testing.T) {
	left, right := readText("treb7uchet")

	if left != 7 {
		t.Error("left value, excepted 7 got ", left)
	}

	if right != 7 {
		t.Error("left value, excepted 7 got ", right)
	}
}
