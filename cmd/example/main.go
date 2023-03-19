package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"log"
	lab2 "github.com/MytsV/kpi-architecture-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Path to an input file")
	outputFile      = flag.String("o", "", "Path to an output file")
)

func getFile(path string) (file *os.File) {
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	flag.Parse()
	if *inputExpression != "" && *inputFile != "" {
		log.Fatal("Multiple input flags specified")
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
		log.Fatal("No input flags specified")
	}

	if *outputFile != "" {
		file := getFile(*outputFile)
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}
	err := handler.Compute()
	if err != nil {
		log.Fatal(err) //Outputs to stderr
	}
}
