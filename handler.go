package lab2

import (
	"io"
)

// The ComputeHandler structbe constructed with input io.Reader and output io.Writer.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Its Compute() method should read the expression from input and write the computed result to the output.
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
