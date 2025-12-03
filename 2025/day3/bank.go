package main

import (
	"math"
)

type bank struct {
	allCells []int
}

func (b *bank) totalJoltage(cellCount int) uint64 {
	activeCells := make([]int, 0)
	cellCandidates := b.allCells

	for i := 0; i < cellCount; i++ {
		candidateCount := len(cellCandidates)
		activeCount := len(activeCells)
		if cellCount-activeCount == candidateCount {
			activeCells = append(activeCells, cellCandidates...)
			break
		} else {
			nextHighestJoltage := 0
			start := 0
			extraCandidates := candidateCount - (cellCount - activeCount)
			for j, c := range cellCandidates {
				if j > extraCandidates {
					break
				}

				if nextHighestJoltage < c {
					nextHighestJoltage = c
					start = j + 1
				}
			}
			activeCells = append(activeCells, nextHighestJoltage)
			cellCandidates = cellCandidates[start:]
		}
	}

	var total uint64
	for i, joltage := range activeCells {
		total += uint64(joltage) * uint64(math.Pow10((cellCount-1)-i))
	}
	return total
}

func newBank(cells []int) *bank {
	return &bank{
		allCells: cells,
	}
}
