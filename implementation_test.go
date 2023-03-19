package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string // test case name
	input          string // function input
	expectedResult string // expected outcome
	hasError       bool   // a flag to check that an error is expected
}

// Define a slice of testCase as test table
var testTable = []testCase{
	//simple input (2 - 3 operands)
	{
		name:           "Simple input; + operator",
		input:          "4 0 4 + +",
		expectedResult: "8",
		hasError:       false,
	},
	{
		name:           "Simple input; - operator",
		input:          "104 506 -",
		expectedResult: "-402",
		hasError:       false,
	},
	{
		name:           "Simple input; * operator",
		input:          "3 4 *",
		expectedResult: "12",
		hasError:       false,
	},
	{
		name:           "Simple input; / operator",
		input:          "6 -3 /",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "Simple input; ^ operator",
		input:          "11 3 ^",
		expectedResult: "1331",
		hasError:       false,
	},
	//complex input
	{
		name:           "Complex input:0",
		input:          "4 2 - 3 * 5 + -1 * -11 / 0 ^",
		expectedResult: "1",
		hasError:       false,
	},
	{
		name:           "Complex input:1",
		input:          "2 3 ^ 4 6 - + 6 * 3 / 2 ^",
		expectedResult: "144",
		hasError:       false,
	},
	//school math restrictions
	{
		name:           "Division by zero",
		input:          "1 0 /",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "Even root of the nagative number",
		input:          "-1 0.5 ^",
		expectedResult: "",
		hasError:       true,
	},
	//float
	{
		name:           "Calculations with float numbers",
		input:          "2.7 -0.3 - 0.1 *",
		expectedResult: "0.3",
		hasError:       false,
	},

	//input ambiguities
	{
		name:           "Not enough operands in the stack",
		input:          "4 0 4 + * -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "Not enough operators in the stack",
		input:          "1 2 3 -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "No operands and operators",
		input:          " ",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "Has an extra space between operands",
		input:          "2 6     2 - -",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "Has an extra space between operators",
		input:          "2 6 2 -     -",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "A character that is not a number, not an operator:0",
		input:          "4! 3 -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "A character that is not a number, not an operator:1",
		input:          "40e 39 %",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "A character that is not a number, not an operator:2",
		input:          "Slava Ysu! Naviky slava!",
		expectedResult: "",
		hasError:       true,
	},
}

func TestEvaluatePostfix(t *testing.T) {
	for _, test := range testTable {

		actual, err := EvaluatePostfix(test.input)
		assert.Equal(t, test.expectedResult, actual, test.name)

		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}

func ExampleEvaluatePostfix() {
	exp1, _ := EvaluatePostfix("9 7 - 2 / ")
	fmt.Printf("exp1: %v\n", exp1)

	exp2, _ := EvaluatePostfix("3 -4 * 6 +")
	fmt.Printf("exp2: %v\n", exp2)

	exp3, _ := EvaluatePostfix("11 2 ^")
	fmt.Printf("exp3: %v\n", exp3)

	// Output:
	// exp1: 1
	// exp2: -6
	// exp3: 121

}
