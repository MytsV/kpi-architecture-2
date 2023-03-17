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

var operations = map[string]func(float64, float64) (float64, error){
	"+": func(a, b float64) (float64, error) {
		return a + b, nil
	},
	"-": func(a, b float64) (float64, error) {
		return b - a, nil
	},
	"*": func(a, b float64) (float64, error) {
		return a * b, nil
	},
	"/": func(a, b float64) (float64, error) {
		if a == 0 {
			return math.NaN(), fmt.Errorf("zero_division")
		}
		return b / a, nil
	},
	"^": func(a, b float64) (float64, error) {
		value := math.Pow(b, a)
		if (math.IsNaN(value)) {
			return value, fmt.Errorf("imaginary_root")
		}
		return value, nil
	},
}

func FloatToString(x float64) string {
	str := fmt.Sprintf("%.6f", x)
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}

// TODO: document this function.
// EvaluatePostfix...
func EvaluatePostfix(input string) (string, error) {
	values := strings.Split(input, " ")
	var stack Stack
	for _, value := range values {
		operation, ok := operations[value]
		if ok {
			//TODO: handle empty stack error
			a, _ := stack.Pop()
			b, _ := stack.Pop()

			//TODO: handle parsing error
			aNumber, _ := strconv.ParseFloat(a, 64)
			bNumber, _ := strconv.ParseFloat(b, 64)

			result, error := operation(aNumber, bNumber)
			if error != nil {
				return "", error
			}
			stack.Push(FloatToString(result))
		} else {
			//TODO: check value?
			stack.Push(value)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("expression_incorrect")
	} else {
		return stack[0], nil
	}
}
