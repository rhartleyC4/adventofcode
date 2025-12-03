package main

import "fmt"

type JoltageCalculator struct {
	bankJolts map[int]int
	nextBank  int
}

type cell struct {
	joltage int
	index   int
}

func (jc *JoltageCalculator) AddBank(input string) (int, error) {
	// test the length
	cell1 := cell{joltage: -1, index: -1}
	cell2 := cell{joltage: -1, index: -1}
	for i, j := range input {
		newJoltage, isNumeric := DigitRuneToInt(j)
		if !isNumeric {
			return 0, fmt.Errorf("invalid character '%c'", j)
		}
		if newJoltage > cell1.joltage && newJoltage > cell2.joltage {
			if i == len(input)-1 {
				// replace the one with the lower value
				if cell1.joltage < cell2.joltage {
					cell1.joltage = newJoltage
					cell1.index = i
				} else {
					cell2.joltage = newJoltage
					cell2.index = i
				}
			} else {
				// reset only one cell
				cell1.joltage = newJoltage
				cell1.index = i
				cell2.joltage = -1
				cell2.index = -1
			}
		} else if newJoltage > cell1.joltage {
			cell1.joltage = newJoltage
			cell1.index = i
		} else if newJoltage > cell2.joltage {
			cell2.joltage = newJoltage
			cell2.index = i
		}
	}

	var totalJolts int
	if cell1.index < cell2.index {
		totalJolts = cell1.joltage*10 + cell2.joltage
	} else {
		totalJolts = cell2.joltage*10 + cell1.joltage
	}

	jc.bankJolts[jc.nextBank] = totalJolts
	jc.nextBank++

	return totalJolts, nil
}

func (jc *JoltageCalculator) Reset() {
	jc.bankJolts = make(map[int]int)
	jc.nextBank = 1
}

func (jc *JoltageCalculator) TotalJolts() int {
	var total int
	for _, jolt := range jc.bankJolts {
		total += jolt
	}
	return total
}

func NewJoltageCalculator() *JoltageCalculator {
	return &JoltageCalculator{
		bankJolts: make(map[int]int),
		nextBank:  1,
	}
}
