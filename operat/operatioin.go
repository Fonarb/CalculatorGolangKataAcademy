package operat

import (
	"fmt"
	"os"
)

func GetResult(x int, s string, y int) int {
	var result int
	switch s {
	case "*":
		result = x * y
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "/":
		if y <= 0 {
			fmt.Println("делитель должен быть > 0")
			os.Exit(1)
		} else {
			result = x / y
		}
	default:
		fmt.Println("Неверное выражение")
		os.Exit(1)
	}
	return result
}
