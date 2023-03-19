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
		name           string
		input          string
		expectedOutput string
		expectedError  error
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

	for _, test := range testTable {
		input := strings.NewReader(test.input)
		output := new(bytes.Buffer)
		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		outputString := output.String()

		assert.Equal(t, test.expectedOutput, outputString, test.name)
		assert.Equal(t, test.expectedError, err, test.name)
	}
}
