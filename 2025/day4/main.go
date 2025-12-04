package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	stock, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("failed to read file %s contents: %v", *filePath, err)
	}

	storage, err := NewStorage(string(stock))
	if err != nil {
		log.Fatalf("failed to create storage for file %s: %v", *filePath, err)
	}

	total := 0
	for {
		rollsToRemove := storage.ForkLiftAccessibleItems(true)
		if rollsToRemove == 0 {
			break
		}
		total += rollsToRemove
	}

	log.Printf("Number of rolls accessible: %d", total)
}
