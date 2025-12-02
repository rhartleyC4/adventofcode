package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "path to input file")
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

	content, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("failed to read file %s contents: %v", *filePath, err)
	}

	inputRanges := strings.Split(string(content), ",")
	var result []uint64
	for _, inputRange := range inputRanges {
		result = append(result, InvalidIDLookup(inputRange)...)
	}

	sum := uint64(0)
	for _, id := range result {
		sum += id
	}
	log.Printf("sum: %d", sum)
}
