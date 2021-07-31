package hackerrank

import (
	"bufio"
	"github.com/clD11/go-katas/src/testsupport"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestHackerrankString_YES(t *testing.T) {
	input := "hereiamstackerrank"
	actual := HackerrankString(input)
	testsupport.AssertThatString(t, actual, "YES")
}

func TestHackerrankString_NO(t *testing.T) {
	input := "hackerworld"
	actual := HackerrankString(input)
	testsupport.AssertThatString(t, actual, "NO")
}

func TestHackerrankString_Longer_NO(t *testing.T) {
	input := "rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt"
	actual := HackerrankString(input)
	testsupport.AssertThatString(t, actual, "NO")
}

func TestHackerrankString_Longest_NO(t *testing.T) {
	input := "rhackerank"
	actual := HackerrankString(input)
	testsupport.AssertThatString(t, actual, "NO")
}

func TestHackerrank_EdgeCases_Short(t *testing.T) {
	input := readData(t, "./testdata/hackerrank_string_short.txt", "./testdata/hackerrank_string_short_ans.txt")
	for _, in := range input {
		actual := HackerrankString(in.input)
		if actual != in.expected {
			log.Println(in.input)
		}
		testsupport.AssertThatString(t, actual, in.expected)
	}
}

func TestHackerrank_EdgeCases(t *testing.T) {
	input := readData(t, "./testdata/hackerrank_string.txt", "./testdata/hackerrank_string_ans.txt")
	for _, in := range input {
		actual := HackerrankString(in.input)
		if actual != in.expected {
			log.Println(in.input)
		}
		testsupport.AssertThatString(t, actual, in.expected)
	}
}

type line struct {
	input    string
	expected string
}

func readData(t *testing.T, input, answers string) []line {
	t.Helper()

	inputFile, _ := os.Open(input)
	defer inputFile.Close()

	answerFile, _ := os.Open(answers)
	defer answerFile.Close()

	inputScanner := bufio.NewScanner(inputFile)
	answerScanner := bufio.NewScanner(answerFile)

	inputScanner.Scan()
	size, _ := strconv.Atoi(inputScanner.Text())
	lines := make([]line, 0, size)

	for i := 0; i < size; i++ {
		inputScanner.Scan()
		answerScanner.Scan()
		in := line{
			input:    inputScanner.Text(),
			expected: answerScanner.Text(),
		}
		lines = append(lines, in)
	}

	return lines
}
