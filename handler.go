package lab2

import (
	"io"
)

// The ComputeHandler struct constructed defines input and output
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute() method should reads the expression from the input, computes it, and write the result to the output.
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
	_, err = ch.Output.Write([]byte(result + "\n"))
	if err != nil {
		return err
	}

	return nil
}
