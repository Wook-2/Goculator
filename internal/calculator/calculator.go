package calculator

import (
	"fmt"
	"math"
	"strconv"

	"github.com/j1mmyson/Goculator/internal/check"
)

//Calculate calculate input
func Calculate(input string, result *float64, storage *[]float64) {
	switch check.IsStartWithOp(input) {
	case true:
		*result = cal(input, *result)
	case false:
		*result, *storage = saveAndCal(input, *result, *storage)
	}
}

func cal(input string, result float64) float64 {
	num, _ := strconv.ParseFloat(input[1:len(input)-1], 64)
	fmt.Println("input", input)
	fmt.Println(result, num)
	switch op := string(input[0]); op {
	case "+":
		result = result + num
	case "-":
		result = result - num
	case "*":
		result = result * num
	case "/":
		result = result / num
	// case "%":
	// 	result = int(num1) % int(num2)
	// 	fmt.Printf("Detected operator is %s\n", op)
	case "^":
		result = math.Pow(result, num)
	}

	return result
}

func saveAndCal(input string, result float64, storage []float64) (float64, []float64) {
	num := result
	storage = append(storage, num)
	result = 0
	if i := check.Operator(input); i == -1 {
		fmt.Println("unvalid input(no operator)")
		// input이 올바르지 않을때 result=0이 그대로 반환된다.
	} else {
		num1, _ := strconv.ParseFloat(input[:i], 64)
		num2, _ := strconv.ParseFloat(input[i+1:len(input)-1], 64)
		switch op := string(input[i]); op {
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
	}

	return result, storage
}
