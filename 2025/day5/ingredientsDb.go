package main

import (
	"fmt"
	"strconv"
	"strings"
)

type FreshRange struct {
	Min uint64
	Max uint64
}

type IngredientsDb struct {
	freshTable []FreshRange
	fresh      []int
	spoiled    []int
}

func (db *IngredientsDb) AddFresh(input string) (*FreshRange, error) {
	values := strings.Split(input, "-")
	parsedMin, err := strconv.ParseUint(values[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing min value: %w", err)
	}
	parsedMax, err := strconv.ParseUint(values[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing max value: %w", err)
	}
	data := &FreshRange{
		Min: parsedMin,
		Max: parsedMax,
	}
	db.freshTable = append(db.freshTable, *data)
	return data, nil
}

func (db *IngredientsDb) CheckIngredient(input string) (bool, error) {
	id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return false, fmt.Errorf("error parsing min value: %w", err)
	}
	for _, data := range db.freshTable {
		if id >= data.Min && id <= data.Max {
			db.fresh = append(db.fresh, int(id))
			return true, nil
		}
	}
	db.spoiled = append(db.spoiled, int(id))
	return false, nil
}

func (db *IngredientsDb) TotalFresh() int {
	return len(db.fresh)
}

func (db *IngredientsDb) Reset() {
	db.freshTable = make([]FreshRange, 0)
	db.fresh = make([]int, 0)
	db.spoiled = make([]int, 0)
}
