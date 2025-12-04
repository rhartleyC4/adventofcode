package main

import (
	"fmt"
	"strings"
)

type Slot struct {
	hasItem bool
}

type Option func(i *Storage)

func WithEmptyRune(opt rune) Option {
	return func(i *Storage) {
		i.emptyRune = opt
	}
}

func WithOccupiedRune(opt rune) Option {
	return func(i *Storage) {
		i.occupiedRune = opt
	}
}

func WithRowSeparator(opt string) Option {
	return func(i *Storage) {
		i.rowSeparator = opt
	}
}

type coordinate struct {
	row int
	col int
}

type gridHelper struct {
	rowSize int
	colSize int
}

func (h gridHelper) adjacentSlots(cell coordinate) []coordinate {
	var adjacentCells []coordinate
	//North West
	if cell.col-1 >= 0 && cell.row-1 >= 0 {
		adjacentCells = append(adjacentCells, coordinate{cell.row - 1, cell.col - 1})
	}
	//North
	if cell.row-1 >= 0 {
		adjacentCells = append(adjacentCells, coordinate{cell.row - 1, cell.col})
	}
	//North East
	if cell.col+1 < h.colSize && cell.row-1 >= 0 {
		adjacentCells = append(adjacentCells, coordinate{cell.row - 1, cell.col + 1})
	}
	// East
	if cell.col+1 < h.colSize {
		adjacentCells = append(adjacentCells, coordinate{cell.row, cell.col + 1})
	}
	//South East
	if cell.col+1 < h.colSize && cell.row+1 < h.rowSize {
		adjacentCells = append(adjacentCells, coordinate{cell.row + 1, cell.col + 1})
	}
	//South
	if cell.row+1 < h.rowSize {
		adjacentCells = append(adjacentCells, coordinate{cell.row + 1, cell.col})
	}
	//South West
	if cell.col-1 >= 0 && cell.row+1 < h.rowSize {
		adjacentCells = append(adjacentCells, coordinate{cell.row + 1, cell.col - 1})
	}
	//West
	if cell.col-1 >= 0 {
		adjacentCells = append(adjacentCells, coordinate{cell.row, cell.col - 1})
	}
	return adjacentCells
}

func newGridHelper(grid [][]Slot) *gridHelper {
	return &gridHelper{
		rowSize: len(grid),
		colSize: len(grid[0]),
	}
}

type Storage struct {
	emptyRune    rune
	occupiedRune rune
	rowSeparator string
	grid         [][]Slot
}

func (s *Storage) ForkLiftAccessibleItems(removeWhenFound bool) int {
	count := 0
	util := newGridHelper(s.grid)
	for rowIndex, row := range s.grid {
		for cellIndex, cell := range row {
			if !cell.hasItem {
				fmt.Print(".")
				continue
			}
			adjacentCells := util.adjacentSlots(coordinate{rowIndex, cellIndex})
			adjacentOccupiedCount := 0
			for _, adjacentCell := range adjacentCells {
				if s.grid[adjacentCell.row][adjacentCell.col].hasItem {
					adjacentOccupiedCount++
				}
			}
			if adjacentOccupiedCount < 4 {
				fmt.Print("x")
				if removeWhenFound {
					s.grid[rowIndex][cellIndex].hasItem = false
				}
				count++
			} else {
				fmt.Print("@")
			}
		}
		fmt.Print("\n")
	}
	return count
}

func NewStorage(stock string, opt ...Option) (*Storage, error) {
	s := &Storage{
		emptyRune:    '.',
		occupiedRune: '@',
		rowSeparator: "\n",
	}
	for _, o := range opt {
		o(s)
	}

	count := strings.Count(stock, string(s.emptyRune))
	count += strings.Count(stock, string(s.occupiedRune))
	count += strings.Count(stock, s.rowSeparator)
	if count != len(stock) {
		return nil, fmt.Errorf("stock has invalid tokens")
	}

	data := strings.Split(stock, s.rowSeparator)
	if len(data) == 1 && data[0] == stock {
		return nil, fmt.Errorf("no data could be parsed from input using separator %q", s.rowSeparator)
	}
	s.grid = make([][]Slot, len(data))
	for i, row := range data {
		s.grid[i] = make([]Slot, len(row))
		for j, cell := range row {
			s.grid[i][j].hasItem = cell == s.occupiedRune
		}
	}

	rowLen := len(s.grid[0])
	for i, row := range s.grid {
		if len(row) != rowLen {
			return nil, fmt.Errorf("grid is jagged at row %d", i)
		}
	}

	return s, nil
}
