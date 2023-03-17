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
		name:           "simple input; + operator",
		input:          "4 0 4 + +",
		expectedResult: "8",
		hasError:       false,
	},
	{
		name:           "simple input; - operator",
		input:          "104 506 -",
		expectedResult: "-402",
		hasError:       false,
	},
	{
		name:           "simple input; * operator",
		input:          "3 4 *",
		expectedResult: "12",
		hasError:       false,
	},
	{
		name:           "simple input; / operator",
		input:          "6 -3 /",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "simple input; ^ operator",
		input:          "11 3 ^",
		expectedResult: "1331",
		hasError:       false,
	},
	//complex input
	{
		name:           "complex input:0",
		input:          "4 2 - 3 * 5 + -1 * -11 / 0 ^",
		expectedResult: "1",
		hasError:       false,
	},
	{
		name:           "complex input:1",
		input:          "2 3 ^ 4 6 - + 6 * 3 / 2 ^",
		expectedResult: "144",
		hasError:       false,
	},
	//school math restrictions
	{
		name:           "division by zero",
		input:          "1 0 /",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "even root of the nagative number",
		input:          "-1 0.5 ^",
		expectedResult: "",
		hasError:       true,
	},
	//float
	{
		name:           "calculations with float numbers",
		input:          "2.7 -0.3 - 0.1 *",
		expectedResult: "",
		hasError:       false,
	},

	//input ambiguities
	{
		name:           "not enough operands in the stack",
		input:          "4 0 4 + * -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "not enough operators in the stack",
		input:          "1 2 3 -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "no operands and operators",
		input:          " ",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "has an extra space between operands",
		input:          "2 6     2 - -",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "has an extra space between operators",
		input:          "2 6 2 -     -",
		expectedResult: "-2",
		hasError:       false,
	},
	{
		name:           "has an extra space between operators",
		input:          "2 6 2 -     -",
		expectedResult: "-2",
		hasError:       true,
	},
	{
		name:           "a character that is not a number, not an operator:0",
		input:          "4! 3 -",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "a character that is not a number, not an operator:1",
		input:          "40e 39 %",
		expectedResult: "",
		hasError:       true,
	},
	{
		name:           "a character that is not a number, not an operator:2",
		input:          "slava Ysu! naviky slava!",
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
	res, _ := EvaluatePostfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 4
}
