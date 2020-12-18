package display

import "fmt"

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
