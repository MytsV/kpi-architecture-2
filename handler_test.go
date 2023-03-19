package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputerHandler(t *testing.T) {
	type testCase struct {
		name           string // test case name
		input          string // input
		expectedOutput string // expected outcome
		hasError       bool   // a flag to check that an error is expected
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

		if v.hasError {
			assert.NotNil(t, err, v.name)
		} else {
			assert.Nil(t, err, v.name)
		}
	}
}
