package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
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
		if math.IsNaN(value) {
			return value, fmt.Errorf("imaginary_root")
		}
		return value, nil
	},
}

//floatToString converts real number to string, excluding trailing zeros.
func floatToString(x float64) string {
	str := fmt.Sprintf("%.6f", x)
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}

/*
EvaluatePostfix computes and returns the result of a mathematical expression in postfix notation.
If expression is incorrect, error is returned.
Functions supports only "+", "-", "*", "/" and "^" operations. Operands must be real numbers.
*/
func EvaluatePostfix(input string) (string, error) {
	values := strings.Fields(input)
	var stack stack
	for _, value := range values {
		operation, ok := operations[value]
		if ok {
			a, success := stack.pop()
			if !success {
				return "", fmt.Errorf("expression_incorrect")
			}
			b, success := stack.pop()
			if !success {
				return "", fmt.Errorf("expression_incorrect")
			}

			aNumber, error := strconv.ParseFloat(a, 64)
			bNumber, error := strconv.ParseFloat(b, 64)
			if error != nil {
				return "", fmt.Errorf("invalid_operand")
			}

			result, error := operation(aNumber, bNumber)
			if error != nil {
				return "", error
			}
			stack.push(floatToString(result))
		} else {
			_, error := strconv.ParseFloat(value, 64)
			if error != nil {
				return "", fmt.Errorf("invalid_operand")
			}
			stack.push(value)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("expression_incorrect")
	} else {
		return stack[0], nil
	}
}
