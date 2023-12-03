package share

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestReadFile(t *testing.T) {
	envFile, _ := godotenv.Read("../.env")

	scanner, err := ReadFile(envFile["TEST_READ_FILE_PATH"])

	if err != nil {
		t.Error(err)
		return
	}

	scanner.Scan()

	if scanner.Text() != "twovgtprdzcjjzkq3ffsbcblnpq" {
		t.Error("First line should be twovgtprdzcjjzkq3ffsbcblnpq")
	}
}
