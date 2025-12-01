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
	minValue := flag.Int("min", 0, "minimum value for tumbler range")
	maxValue := flag.Int("max", 99, "maximum value for tumbler range")
	flag.Parse()

	if *filePath == "" {
		if flag.NArg() > 0 {
			*filePath = flag.Arg(0)
		}
	}

	if *filePath == "" {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s -file <path> [-min N] [-max M]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *minValue > *maxValue {
		log.Fatalf("min (%d) greater than max (%d)", *minValue, *maxValue)
	}

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", *filePath, err)
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 1024)
	scanner.Buffer(buf, 1024*1024) // up to 1MB line

	tumbler := NewElfTumbler(*minValue, *maxValue)
	lineNum := 0
	applied := 0

	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		adjustment := strings.TrimSpace(raw)
		if adjustment == "" || strings.HasPrefix(adjustment, "#") { // skip empty or comment lines
			continue
		}
		if len(adjustment) < 2 { // must have direction + number
			log.Fatalf("invalid adjustment on line %d: %q", lineNum, adjustment)
		}
		if err := tumbler.Adjust(adjustment); err != nil {
			log.Fatalf("error applying adjustment on line %d (%q): %v", lineNum, adjustment, err)
		} else {
			log.Printf("adjusted tumbler at line %d: %s", lineNum, adjustment)
		}
		applied++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	log.Printf("Processed %d adjustments from %s\n", applied, *filePath)
	log.Printf("Password: %d\n", tumbler.GetPassword())
}
