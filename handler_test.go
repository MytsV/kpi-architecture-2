package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputerHandler(t *testing.T) {
	type testCase struct {
		name           string // test case name
		input          string // input
		expectedOutput string // expected outcome
		expectedError  error  // expected error message
	}

	var testTable = []testCase{
		{
			name:           "Simple input",
			input:          "4 0 4 + +",
			expectedOutput: "8",
			expectedError:  nil,
		},
		{
			name:           "Error: expression_incorrect",
			input:          "4 0 4",
			expectedOutput: "",
			expectedError:  errors.New("expression_incorrect"),
		},
		{
			name:           "Error: zero_division",
			input:          "12 78 3 0 / + *",
			expectedOutput: "",
			expectedError:  errors.New("zero_division"),
		},
	}

	for _, v := range testTable {

		input := strings.NewReader(v.input)
		output := new(bytes.Buffer)

		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		outputString := output.String()

		assert.Equal(t, v.expectedOutput, outputString, v.name)
		assert.Equal(t, v.expectedError, err, v.name)
	}
}
