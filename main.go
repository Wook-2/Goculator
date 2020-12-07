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
func SaveAndCal(input string, result *float64, storage *[]float64) {
	num := *result
	*storage = append(*storage, num)
	*result = 0
	if i := strings.IndexAny(input, o); i == -1 {
		fmt.Println("unvalid input(no operator)")
		// input이 올바르지 않을때 result=0이 그대로 반환된다.
	} else {
		num1, _ := strconv.ParseFloat(input[:i], 64)
		num2, _ := strconv.ParseFloat(input[i+1:len(input)-1], 64)
		switch op := string(input[i]); op {
		case "+":
			*result = num1 + num2
		case "-":
			*result = num1 - num2
		case "*":
			*result = num1 * num2
		case "/":
			*result = num1 / num2
		case "^":
			*result = math.Pow(num1, num2)
		}
	}

	return
}

//Display will display result of calculate and storage space
func Display(storage []float64, result float64, input string) {

	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================")
	fmt.Println("|             GOCULATOR             |")
	fmt.Println("=====================================")
	fmt.Println("|storage|                           |")
	fmt.Println(storage)
	fmt.Println("=====================================")
	fmt.Println("|    Load from storage: type s[n]   |")
	fmt.Println("|          (ex. s[0] + 20)          |")
	fmt.Println("=====================================")
	fmt.Println("|Calculator|")
	fmt.Printf("%s", input)
	fmt.Println("=", result)
	fmt.Println("(enter 'x' to exit)")

	return
}

// CheckCallStorage convert input value if it contains storage
func CheckCallStorage(input string, storage []float64) string {
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

func main() {

	kbReader := bufio.NewReader(os.Stdin)
	storage := make([]float64, 0)
	// st_ind := 0
	var op string
	var num1, num2 float64
	var result float64
	var input string
	Display(storage, result, input)
	input, _ = kbReader.ReadString('\n')        // 엔터키가 나올때까지 입력을 받음.
	input = strings.Replace(input, " ", "", -1) // input에서 모든 공백 제거
	input = CheckCallStorage(input, storage)    // 저장된 값을 불러오는지 체크하고 인풋을 그에맞게 변경해줌.

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
		Display(storage, result, input)
		for {
			input, _ = kbReader.ReadString('\n')
			input = strings.Replace(input, " ", "", -1)
			input = CheckCallStorage(input, storage)
			if so := StartsWithOp(input); so {
				Calculate(input, &result)
				Display(storage, result, input)
			} else {
				SaveAndCal(input, &result, &storage)
				Display(storage, result, input)
			}
			if string(input[0]) == "x" {
				fmt.Print("\033[H\033[2J")
				fmt.Println("Turn off")
				break
			}
		}
	}

}
