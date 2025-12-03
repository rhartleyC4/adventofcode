package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

	calculator := NewJoltageCalculator(12)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		bank := strings.TrimSpace(raw)
		if bank == "" || strings.HasPrefix(bank, "#") { // skip empty or comment lines
			continue
		}
		if jolts, err := calculator.AddBank(bank); err != nil {
			log.Fatalf("error adding bank on line %d (%q): %v", lineNum, bank, err)
		} else {
			log.Printf("added %d jolts from bank at line %d: %s", jolts, lineNum, bank)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	log.Printf("Total Jolts: %d\n", calculator.TotalJolts())
}
