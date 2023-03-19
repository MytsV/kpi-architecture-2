package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/MytsV/kpi-architecture-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Path to an input file")
	outputFile      = flag.String("o", "", "Path to an output file")
)

func getFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	flag.Parse()
	if *inputExpression != "" && *inputFile != "" {
		panic("Multiple input flags specified")
	}

	var (
		reader io.Reader
		writer io.Writer
	)

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file := getFile(*inputFile)
		defer file.Close()
		reader = file
	} else {
		panic("No input flags specified")
	}

	if *outputFile != "" {
		file := getFile(*outputFile)
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	buf := make([]byte, 1024)

	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	result, err := lab2.EvaluatePostfix(string(buf[:n]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating postfix: %v\n", err)
		os.Exit(1)
	}

	_, err = writer.Write([]byte(result))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	// handler := &lab2.ComputeHandler{
	// 	Input: reader,
	// 	Output: writer,
	// }
	//handler.Compute()
}
