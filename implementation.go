package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// TODO: document this function.
// EvaluatePostfix...
func EvaluatePostfix(input string) (string, error) {
	values := strings.Split(input, " ")
	stack := make([]string, 0)
	for _, value := range values {
		if value == "+" || value == "-" || value == "*" || value == "/" || value == "^" {
			index := len(stack) - 1
			a, _ := strconv.Atoi(stack[index])
			stack = stack[:index]
			index = len(stack) - 1
			b, _ := strconv.Atoi(stack[index])
			stack = stack[:index]
			var result int
			if value == "+" {
				result = a + b
			}
			if value == "*" {
				result = a * b
			}
			if value == "/" {
				result = b / a
			}
			if value == "-" {
				result = b - a
			}
			if value == "^" {
				result = int(math.Pow(float64(b), float64(a)))
			}
			stack = append(stack, strconv.Itoa(result))
		} else {
			stack = append(stack, value)
		}
	}
	for _, value := range stack {
		return value, nil
	}
	return "TODO", fmt.Errorf("TODO")
}
