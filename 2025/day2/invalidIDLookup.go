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
		if len(candidate)%2 != 0 {
			continue
		}
		offset := len(candidate) / 2
		invalidID := true
		for i := 0; i < offset; i++ {
			if candidate[i] != candidate[offset+i] {
				invalidID = false
				break
			}
		}
		if invalidID {
			ids = append(ids, i)
		}
	}
	return ids
}
