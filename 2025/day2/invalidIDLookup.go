package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func InvalidIDLookup(r string) []uint64 {
	var ids = make([]uint64, 0)
	rangeSplit := strings.Split(r, "-")
	if len(rangeSplit) != 2 {
		log.Fatalf("Invalid ranges: %s", r)
		return ids
	}
	start, err := strconv.ParseUint(rangeSplit[0], 10, 64)
	if err != nil {
		log.Fatalf("Invalid start: %s", rangeSplit[0])
		return ids
	}
	end, err := strconv.ParseUint(rangeSplit[1], 10, 64)
	if err != nil {
		log.Fatalf("Invalid end: %s", rangeSplit[1])
		return ids
	}
	for i := start; i <= end; i++ {
		candidate := fmt.Sprintf("%d", i)
		candidateLength := len(candidate)
		if candidateLength <= 1 {
			continue
		}
		flagInvalid := false
		for j := candidateLength / 2; j >= 1; j-- {
			testLength := GCD(candidateLength, j)
			parts := candidateLength / testLength
			for k := 0; k < parts-1; k++ {
				leftStart := k * testLength
				leftEnd := leftStart + testLength
				rightEnd := leftEnd + testLength
				flagInvalid = candidate[leftStart:leftEnd] == candidate[leftEnd:rightEnd]
				if !flagInvalid {
					break
				}
			}
			if flagInvalid {
				break
			}
		}
		if flagInvalid {
			ids = append(ids, i)
		}
	}
	return ids
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
