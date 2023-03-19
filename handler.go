package lab2

import (
	"io"
)

// ComputeHandler struct defines input and output via default interfaces
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute method reads the expression from the input, evaluates it, and writes the result to the output.
func (ch *ComputeHandler) Compute() error {
	inputBuffer := make([]byte, 256)
	n, err := ch.Input.Read(inputBuffer)
	if err != nil {
		return err
	}
	result, err := EvaluatePostfix(string(inputBuffer[:n]))
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}
