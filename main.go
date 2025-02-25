package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func main() {
	printLineNumbers := false
	flag.BoolVar(&printLineNumbers, "n", false, "print line numbers")
	flag.Parse()
	
	flags := flag.Args()

	if len(flags) == 0 || flags[0] == "-"  {
		read(os.Stdin, printLineNumbers)
		return
	}

	if slices.Contains(flag.Args(), "-n") {
		printLineNumbers = true
	}

	for _, fileName := range flag.Args() {
		if fileName == "-n" {
			continue
		}

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("opening file %s: %v", fileName, err)
		}

		err = read(f, printLineNumbers)
		if err != nil {
			log.Printf("reading file %s: %v", fileName, err)
			continue
		}
	}
}

func read(r io.Reader, printNum bool) error {
	reader := bufio.NewScanner(r)
	lineNumber := 1

	for reader.Scan() {
		if printNum {
			fmt.Printf("%d %s\n", lineNumber, reader.Text())
			lineNumber++

			continue
		}
		fmt.Println(reader.Text())
	}

	if err := reader.Err(); err != nil {
		return err
	}

	return nil
}
