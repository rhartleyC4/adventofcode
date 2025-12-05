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

	db := IngredientsDb{}
	lineNum := 0
	processFresh := true
	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		entry := strings.TrimSpace(raw)
		if entry == "" {
			processFresh = false
			continue
		}
		if processFresh {
			_, err := db.AddFresh(entry)
			if err != nil {
				log.Fatalf("failed to parse fresh range %q on line %d: %v", entry, lineNum, err)
			}
		} else {
			_, err := db.CheckIngredient(entry)
			if err != nil {
				log.Fatalf("failed to check ingredient on %q on line %d: %v", entry, lineNum, err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	log.Printf("Total fresh ingredients: %d\n", db.TotalFresh())
	log.Printf("Total fresh candidates: %d\n", db.TotalConsideredFresh())
}
