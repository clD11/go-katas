package gokatas

import (
	"strings"
	"strconv"
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')	
	input, _ := reader.ReadString('\n')
	fmt.Print(Sum(input))
}

func Sum(input string) int64 {	
	summable := strings.Split(input, " ")

	var sum int64
	for _, num := range summable {
		if n, err := strconv.ParseInt(num, 10, 64); err == nil {
			sum = sum + n
		}		
	}

	return sum
}