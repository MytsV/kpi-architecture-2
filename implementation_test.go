package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluatePostfix(t *testing.T) {
	res, err := EvaluatePostfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "11", res)
	}
	res, err = EvaluatePostfix("7 2 ^ 25 10 5 / + * 13 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "1310", res)
	}
	res, err = EvaluatePostfix("3 5 2 ^ * 15 / 5 2 2 ^ - -")
	if assert.Nil(t, err) {
		assert.Equal(t, "4", res)
	}
	res, err = EvaluatePostfix("18 3 / 2 ^ 13 7 + 5 2 ^ * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "536", res)
	}
}

func ExampleEvaluatePostfix() {
	res, _ := EvaluatePostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
