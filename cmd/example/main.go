package main

import (
	"flag"
	"os"
	"io"
	"strings"
	lab2 "github.com/MytsV/kpi-architecture-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Path to an input file")
	outputFile = flag.String("o", "", "Path to an output file")
)

func getFile(path string) (file *os.File){
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if (err != nil) {
		panic(err)
	}
	return
}

func main() {
	flag.Parse()
	if (*inputExpression != "" && *inputFile != "") {
		panic("Multiple input flags specified")
	}

	var (
		reader io.Reader
		writer io.Writer
	)

	if (*inputExpression != "") {
		reader = strings.NewReader(*inputExpression)
	} else if (*inputFile != "") {
		file := getFile(*inputFile)
		defer file.Close()
		reader = file
	} else {
		panic("No input flags specified")
	}

	if (*outputFile != "") {
		file := getFile(*outputFile)
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	buf := make([]byte, 1024)
	//TODO: handle errors on working with io
	n, _ := reader.Read(buf)
	result, _ := lab2.EvaluatePostfix(string(buf[:n]))
	writer.Write([]byte(result))

	// handler := &lab2.ComputeHandler{
	// 	Input: reader,
	// 	Output: writer,
	// }
	//handler.Compute()
}
