package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var o string = "+-*/%^"

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

//Calculate calculate input
func Calculate(input string, result *float64, storage *[]float64) {
	switch IsStartWithOp(input) {
	case true:
		Cal(input, &result)
	case false:
		SaveAndCal(input, &result, &storage)
	}
}

// Cal : 연산자로 시작하는 수식을 계산한다
func Cal(input string, result **float64) {
	num, _ := strconv.ParseFloat(input[1:len(input)-1], 64)
	fmt.Println("input", input)
	fmt.Println(**result, num)
	switch op := string(input[0]); op {
	case "+":
		**result = **result + num
	case "-":
		**result = **result - num
	case "*":
		**result = **result * num
	case "/":
		**result = **result / num
	// case "%":
	// 	result = int(num1) % int(num2)
	// 	fmt.Printf("Detected operator is %s\n", op)
	case "^":
		**result = math.Pow(**result, num)
	}

	return
}

// SaveAndCal save the value and start new calculation
func SaveAndCal(input string, result **float64, storage **[]float64) {
	num := **result
	**storage = append(**storage, num)
	**result = 0
	if i := strings.IndexAny(input, o); i == -1 {
		fmt.Println("unvalid input(no operator)")
		// input이 올바르지 않을때 result=0이 그대로 반환된다.
	} else {
		num1, _ := strconv.ParseFloat(input[:i], 64)
		num2, _ := strconv.ParseFloat(input[i+1:len(input)-1], 64)
		switch op := string(input[i]); op {
		case "+":
			**result = num1 + num2
		case "-":
			**result = num1 - num2
		case "*":
			**result = num1 * num2
		case "/":
			**result = num1 / num2
		case "^":
			**result = math.Pow(num1, num2)
		}
	}

	return
}

// ShowEndPage show end page and turn off the program
func ShowEndPage() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("=====================================")
	fmt.Println("|              TURN OFF             |")
	fmt.Println("=====================================")
}

//Display will display result of Cal and storage space
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

// ReadInput read keboard input and remove empty space.
func ReadInput() (input string) {
	kbReader := bufio.NewReader(os.Stdin)
	input, _ = kbReader.ReadString('\n')        // 엔터키가 나올때까지 입력을 받음.
	input = strings.Replace(input, " ", "", -1) // input에서 모든 공백 제거

	return
}

// CheckErrors will check error from input
func CheckErrors() {
	// 	1. operator doesn`t exist.
	//  2. input[0] == "x" -> turn off
	//  3. unvalid input -> continue;
	//  4. ...
	return
}

// Goculator is main function of Goculator
func Goculator() {
	storage := make([]float64, 0)
	var (
		input  string
		result float64
	)

	for {
		Display(storage, result, input)
		input = ReadInput()
		input = CheckCallStorage(input, storage) // 저장된 값을 불러오는지 체크하고 인풋을 그에맞게 변경해줌.
		if strings.HasPrefix(input, "x") {
			ShowEndPage()
			break
		}
		Calculate(input, &result, &storage)
	}
}

func main() {
	Goculator()
	return
}
