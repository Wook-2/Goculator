package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 기능 :
// 사칙연산
// 나머지연산
// 로그연산 (later)
// 제곱연산
// 괄호연산 (later)
// 값 저장, 불러오기
// var op = [6]string{"+", "-", "*", "/", "%", "^"}

var o string = "+-*/%^"

// StartsWithOp check if the input starts with operator
func StartsWithOp(input string) bool {
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

// Calculate : 순차적 진행
func Calculate(input string, result *float64) {
	num, _ := strconv.ParseFloat(input[1:len(input)-1], 64)
	fmt.Println("input", input)
	fmt.Println(*result, num)
	switch op := string(input[0]); op {
	case "+":
		*result = *result + num
	case "-":
		*result = *result - num
	case "*":
		*result = *result * num
	case "/":
		*result = *result / num
	// case "%":
	// 	result = int(num1) % int(num2)
	// 	fmt.Printf("Detected operator is %s\n", op)
	case "^":
		*result = math.Pow(*result, num)
	}
	return
}

// SaveAndCal save the value and start new calculation
func SaveAndCal() {

	return
}

func main() {

	kbReader := bufio.NewReader(os.Stdin)
	var op string
	var num1, num2 float64
	var result float64
	input, _ := kbReader.ReadString('\n')       // 엔터키가 나올때까지 입력을 받음.
	input = strings.Replace(input, " ", "", -1) // input에서 모든 공백 제거

	if i := strings.IndexAny(input, o); i == -1 {
		fmt.Println("unvalid input(no operator)")
	} else {
		num1, _ = strconv.ParseFloat(input[:i], 64)
		num2, _ = strconv.ParseFloat(input[i+1:len(input)-1], 64)
		switch op = string(input[i]); op {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		case "^":
			result = math.Pow(num1, num2)
		}
		fmt.Println(num1, op, num2, "=", result)
		for {
			input, _ = kbReader.ReadString('\n')        // 엔터키가 나올때까지 입력을 받음.
			input = strings.Replace(input, " ", "", -1) // input에서 모든 공백 제거
			if string(input[0]) == "x" {
				break
			}
			Calculate(input, &result)
			fmt.Println(result)
		}

	}

}
