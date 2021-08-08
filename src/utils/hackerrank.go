package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readLines() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lines, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < lines; i++ {
		value := strings.Trim(scanner.Text(), "\n\t")
		fmt.Println(value)
	}
}

func ReadLines(t *testing.T, filepath string, numLines int) []string {
	t.Helper()

	reader, err := os.Open(filepath)
	defer reader.Close()

	if err != nil {
		t.Fatal(err)
	}

	scan := bufio.NewScanner(reader)

	lines := make([]string, 0, numLines)
	for i := 0; i < numLines; i++ {
		scan.Scan()
		text := scan.Text()
		lines = append(lines, text)
	}

	return lines
}
