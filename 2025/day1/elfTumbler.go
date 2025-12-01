package main

import (
	"fmt"
	"strconv"
)

type ElfTumbler struct {
	zeroCount    int
	currentIndex int
	tumblerSize  int
}

func (t *ElfTumbler) GetValue() int {
	return t.currentIndex
}

func (t *ElfTumbler) Adjust(adjustment string) error {
	direction := adjustment[0]
	advanceBy, err := strconv.Atoi(adjustment[1:])
	if err != nil {
		return err
	}
	t.zeroCount += advanceBy / t.tumblerSize
	adjustedAdvanceBy := advanceBy % t.tumblerSize
	switch direction {
	case 'L':
		newIndex := t.currentIndex - adjustedAdvanceBy
		if newIndex < 0 {
			if t.currentIndex != 0 {
				t.zeroCount++
			}
			t.currentIndex = newIndex + t.tumblerSize
		} else {
			t.currentIndex = newIndex
		}
	case 'R':
		newIndex := t.currentIndex + adjustedAdvanceBy
		if newIndex >= t.tumblerSize {
			t.currentIndex = newIndex % t.tumblerSize
			if t.currentIndex != 0 {
				t.zeroCount++
			}
		} else {
			t.currentIndex = newIndex
		}
	default:
		return fmt.Errorf("unknown direction '%c'", direction)
	}

	if t.GetValue() == 0 {
		t.zeroCount++
	}
	return nil
}

func (t *ElfTumbler) GetPassword() int {
	return t.zeroCount
}

func NewElfTumbler(min, max int) *ElfTumbler {
	if min > max {
		panic("min > max")
	}
	size := max - min + 1
	return &ElfTumbler{
		zeroCount:    0,
		currentIndex: size / 2,
		tumblerSize:  size,
	}
}
