package main

import "fmt"

func calculate(condition string, values ...int) int {
	if len(values) == 0 {
		return 0
	}

	switch condition {
	case "*":
		result := 1
		for _, number := range values {
			result *= number
		}
		return result

	case "/":
		result := values[0]
		for _, number := range values[1:] {
			if number == 0 {
				fmt.Println("Error: Division by zero")
				return 0
			}
			result /= number
		}
		return result

	case "-":
		result := values[0]
		for _, number := range values[1:] {
			result -= number
		}
		return result

	default: // "+" case and any other input
		result := 0
		for _, number := range values {
			result += number
		}
		return result
	}
}

