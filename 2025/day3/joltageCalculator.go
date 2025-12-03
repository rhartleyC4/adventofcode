package main

import (
	"fmt"
)

type JoltageCalculator struct {
	banks         []*bank
	activateCells int
}

func (jc *JoltageCalculator) AddBank(input string) (uint64, error) {
	cells := make([]int, len(input))
	for i, j := range input {
		joltage, isNumeric := DigitRuneToInt(j)
		if !isNumeric {
			return 0, fmt.Errorf("invalid character '%c'", j)
		}
		cells[i] = joltage
	}

	b := newBank(cells)
	jc.banks = append(jc.banks, b)

	return b.totalJoltage(jc.activateCells), nil
}

func (jc *JoltageCalculator) Reset() {
	jc.banks = make([]*bank, 0)
}

func (jc *JoltageCalculator) TotalJolts() uint64 {
	var total uint64
	for _, b := range jc.banks {
		total += b.totalJoltage(jc.activateCells)
	}
	return total
}

func NewJoltageCalculator(activateCells int) *JoltageCalculator {
	return &JoltageCalculator{
		banks:         make([]*bank, 0),
		activateCells: activateCells,
	}
}
