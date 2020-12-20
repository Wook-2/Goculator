package goculator

import (
	"github.com/j1mmyson/Goculator/internal/calculator"
	"github.com/j1mmyson/Goculator/internal/check"
	"github.com/j1mmyson/Goculator/internal/display"
)

// Goculator is main function of Goculator
func Goculator() {
	storage := make([]float64, 0)
	var input string
	var result float64

	for {
		display.Display(storage, result, input)
		input = check.ReadInput()
		if check.Exit(input) {
			break
		}
		if check.Errors() {
			break
		}
		input = check.CallStorage(input, storage) // 저장된 값을 불러오는지 체크하고 인풋을 그에맞게 변경해줌.
		calculator.Calculate(input, &result, &storage)
	}
}
