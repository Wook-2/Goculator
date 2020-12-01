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

func main() {

	kbReader := bufio.NewReader(os.Stdin)
	var op string
	var num1, num2 float64
	var result float64
	input, _ := kbReader.ReadString('\n')
	input = strings.Replace(input, " ", "", -1)

	if i := strings.IndexAny(input, o); i == -1 {
		fmt.Println("unvalid input(no operator)")
	} else {
		num1, _ = strconv.ParseFloat(input[:i], 64)
		num2, _ = strconv.ParseFloat(input[i+1:len(input)-1], 64)
		switch op = string(input[i]); op {
		case "+":
			result = num1 + num2
			fmt.Printf("Detected operator is %s\n", op)
		case "-":
			result = num1 - num2
			fmt.Printf("Detected operator is %s\n", op)
		case "*":
			result = num1 * num2
			fmt.Printf("Detected operator is %s\n", op)
		case "/":
			result = num1 / num2
			fmt.Printf("Detected operator is %s\n", op)
		// case "%":
		// 	result = int(num1) % int(num2)
		// 	fmt.Printf("Detected operator is %s\n", op)
		case "^":
			result = math.Pow(num1, num2)
			fmt.Printf("Detected operator is %s\n", op)
		}

		fmt.Println(num1, op, num2, "=", result)
	}

}
