package check

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsStartWithOp check if the input starts with operator
func IsStartWithOp(input string) bool {
	a := strings.HasPrefix(input, "*")
	b := strings.HasPrefix(input, "+")
	c := strings.HasPrefix(input, "-")
	d := strings.HasPrefix(input, "/")
	e := strings.HasPrefix(input, "^")

	if a || b || c || d || e {
		return true
	}

	return false
}

// CallStorage convert input value if it contains storage
func CallStorage(input string, storage []float64) string {
	for {
		ind := strings.Index(input, "s[")
		if ind == -1 {
			break
		}
		num, _ := strconv.ParseInt(input[ind+2:ind+3], 10, 32)
		s := fmt.Sprintf("%f", storage[num])
		input = strings.Replace(input, input[ind:ind+4], s, 1)
	}

	return input
}

// ReadInput read keboard input and remove empty space.
func ReadInput() (input string) {
	kbReader := bufio.NewReader(os.Stdin)
	input, _ = kbReader.ReadString('\n')        // 엔터키가 나올때까지 입력을 받음.
	input = strings.Replace(input, " ", "", -1) // input에서 모든 공백 제거

	return
}

// Operator check if input has an operator
func Operator(input string) int {
	var o string = "+-*/%^"

	return strings.IndexAny(input, o)
}

// Errors will check error from input
func Errors() {
	// 	1. operator doesn`t exist.
	//  2. input[0] == "x" -> turn off
	//  3. unvalid input -> continue;
	//  4. ...
	return
}
