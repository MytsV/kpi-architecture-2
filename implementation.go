package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// TODO: document this function.
// EvaluatePostfix...
func EvaluatePostfix(input string) (string, error) {
	values := strings.Split(input, " ")
	var stack Stack
	for _, value := range values {
		if value == "+" || value == "-" || value == "*" || value == "/" || value == "^" {
			a, _ := stack.Pop()
			b, _ := stack.Pop()

			aNumber, _ := strconv.Atoi(a)
			bNumber, _ := strconv.Atoi(b)

			var result int
			if value == "+" {
				result = aNumber + bNumber
			}
			if value == "*" {
				result = aNumber * bNumber
			}
			if value == "/" {
				result = bNumber / aNumber
			}
			if value == "-" {
				result = bNumber - aNumber
			}
			if value == "^" {
				result = int(math.Pow(float64(bNumber), float64(aNumber)))
			}
			stack.Push(strconv.Itoa(result))
		} else {
			stack.Push(value)
		}
	}
	for _, value := range stack {
		return value, nil
	}
	return "TODO", fmt.Errorf("TODO")
}
