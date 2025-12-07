package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := flag.String("file", "", "path to input file containing adjustments (one per line)")
	flag.Parse()

	if *filePath == "" {
		if flag.NArg() > 0 {
			*filePath = flag.Arg(0)
		}
	}

	if *filePath == "" {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s -file <path>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", *filePath, err)
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 1024)
	scanner.Buffer(buf, 1024*1024) // up to 1MB line

	homework := Homework{}
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		err = homework.AddLine(raw)
		if err != nil {
			log.Fatalf("failed to add line %d: %v", lineNum, err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	solution, err := homework.Solve()
	if err != nil {
		log.Fatalf("failed to solve homework: %v", err)
	}
	log.Printf("Homework solution: %d\n", solution)
}
