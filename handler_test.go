package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputerHandler(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedOutput string
		hasError       bool
	}

	var testTable = []testCase{
		{
			name:           "Simple input",
			input:          "4 0 4 + +",
			expectedOutput: "8",
			hasError:       false,
		},
		{
			name:           "Has an error",
			input:          "4 0 4",
			expectedOutput: "",
			hasError:       true,
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
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
