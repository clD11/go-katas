package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
